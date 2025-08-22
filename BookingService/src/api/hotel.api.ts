import axios from 'axios';
import {serverConfig} from "../config";

export const getAvailableRooms = async (roomCategoryId: number, checkInDate: string, checkOutDate: string) => {
    const response = await axios.post(`${serverConfig.HOTEL_SERVICE_URL}/rooms/available`, {
            roomCategoryId,
            checkInDate,
            checkOutDate
    })
    console.log(response.data.data);
    return response.data.data;
}

export const updateBookingIdToRooms = async (bookingId: number, roomIds: number[]) => {
    const response = await axios.post(`${serverConfig.HOTEL_SERVICE_URL}/rooms/update-booking-id`, {
        bookingId,
        roomIds
    });
    return response.data;
}