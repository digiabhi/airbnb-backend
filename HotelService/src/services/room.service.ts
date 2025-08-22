import {RoomRepository} from "../repositories/room.repository";
import {GetAvailableRoomsDTO, UpdateBookingIdToRoomsDTO} from "../dto/room.dto";

const roomRepository = new RoomRepository();

export async function getAvailableRoomsService(getAvailableRoomsDTO: GetAvailableRoomsDTO) {
    const availableRooms = await roomRepository.findByRoomCategoryIdAndDateRange(getAvailableRoomsDTO.roomCategoryId, new Date(getAvailableRoomsDTO.checkInDate), new Date(getAvailableRoomsDTO.checkOutDate));
    console.log(new Date(getAvailableRoomsDTO.checkInDate), new Date(getAvailableRoomsDTO.checkOutDate))
    return availableRooms;
}

export async function updateBookingIdToRoomsService(updateBookingIdToRoomsDTO: UpdateBookingIdToRoomsDTO) {
    return await roomRepository.updateBookingIdToRooms(updateBookingIdToRoomsDTO.bookingId, updateBookingIdToRoomsDTO.roomIds);
}