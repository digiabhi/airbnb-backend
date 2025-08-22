import {
  confirmBooking,
  createBooking,
  createIdempotencyKey,
  finalizeIdempotencyKey,
  getIdempotencyKeyWithLock,
} from '../repositories/booking.repository';
import { generateIdempotencyKey } from '../utils/generateIdempotencyKey';
import {
  BadRequestError,
  InternalServerError,
  NotFoundError,
} from '../utils/errors/app.error';
import { CreateBookingDTO } from '../dto/booking.dto';
import prismaClient from '../prisma/client';
import { redlock } from '../config/redis.config';
import { serverConfig } from '../config';
import {getAvailableRooms, updateBookingIdToRooms} from "../api/hotel.api";

type AvailableRoom = {
    id: number;
    roomCategoryId: number;
    dateOfAvailability: Date;
}

export async function createBookingService(createBookingDTO: CreateBookingDTO) {
  const ttl = serverConfig.LOCK_TTL; // 5 minutes in milliseconds

    const availableRooms = await getAvailableRooms(createBookingDTO.roomCategoryId, createBookingDTO.checkInDate, createBookingDTO.checkOutDate);

    const checkInDate = new Date(createBookingDTO.checkInDate);
    const checkOutDate = new Date(createBookingDTO.checkOutDate);
    const totalNights = Math.ceil((checkOutDate.getTime() - checkInDate.getTime()) / (1000 * 60 * 60 * 24));

    if (availableRooms.length === 0 || availableRooms.length < totalNights) {
        throw new BadRequestError('No rooms available for the given dates');
    }

    const roomResources = availableRooms.map((room: AvailableRoom) => `room:${room.id}`);
    let lock;

    try {
        lock = await redlock.acquire(roomResources, ttl);
        const booking = await createBooking({
        userId: createBookingDTO.userId,
        hotelId: createBookingDTO.hotelId,
        totalGuests: createBookingDTO.totalGuests,
        bookingAmount: createBookingDTO.bookingAmount,
        checkInDate: new Date(createBookingDTO.checkInDate).toISOString(),
        checkOutDate: new Date(createBookingDTO.checkOutDate).toISOString(),
        roomCategoryId: createBookingDTO.roomCategoryId,
    });

    const idempotencyKey = generateIdempotencyKey();

    await createIdempotencyKey(idempotencyKey, booking.id);
    console.log("idempotencyKey generated");

    await updateBookingIdToRooms(booking.id, availableRooms.map((room: AvailableRoom) => room.id));

    return {
      bookingId: booking.id,
      idempotencyKey: idempotencyKey,
    };
  } catch (error) {
        console.error('Error creating booking:', error);
    throw new InternalServerError(
        `Failed to acquire locks for room resources`
    );
    } finally {
        if (lock) {
            await lock.release().catch((error) => {
                console.error('Error releasing lock:', error);
            });
        }
    }
  // return await redlock.using([bookingResource], ttl, async () => {
  //   const booking = await createBooking({
  //   userId: createBookingDTO.userId,
  //   hotelId: createBookingDTO.hotelId,
  //   totalGuests: createBookingDTO.totalGuests,
  //   bookingAmount: createBookingDTO.bookingAmount,
  // });

  // const idempotencyKey = generateIdempotencyKey();

  // await createIdempotencyKey(idempotencyKey, booking.id);

  // return {
  //   bookingId: booking.id,
  //   idempotencyKey: idempotencyKey,
  // };
  // }
}

export async function confirmBookingService(idempotencyKey: string) {
  return prismaClient.$transaction(async (tx) => {
      const idempotencyKeyData = await getIdempotencyKeyWithLock(
          tx,
          idempotencyKey
      );

      if (!idempotencyKeyData || !idempotencyKeyData.bookingId) {
          throw new NotFoundError('Idempotency key not found');
      }

      if (idempotencyKeyData.finalized) {
          throw new BadRequestError('Idempotency key already finalized');
      }

      const booking = await confirmBooking(tx, idempotencyKeyData.bookingId);
      await finalizeIdempotencyKey(tx, idempotencyKey);

      return booking;
  });
}
