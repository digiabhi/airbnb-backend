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

export async function createBookingService(createBookingDTO: CreateBookingDTO) {
  const ttl = serverConfig.LOCK_TTL; // 5 minutes in milliseconds
  const bookingResource = `hotel:${createBookingDTO.hotelId}`;

  try {
    await redlock.acquire([bookingResource], ttl);  
    const booking = await createBooking({
      userId: createBookingDTO.userId,
      hotelId: createBookingDTO.hotelId,
      totalGuests: createBookingDTO.totalGuests,
      bookingAmount: createBookingDTO.bookingAmount,
    });

    const idempotencyKey = generateIdempotencyKey();

    await createIdempotencyKey(idempotencyKey, booking.id);

    return {
      bookingId: booking.id,
      idempotencyKey: idempotencyKey,
    };
  } catch (error) {
    throw new InternalServerError(
      `Failed to acquire lock for booking resource`
    );
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
  return await prismaClient.$transaction(async (tx) => {
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
