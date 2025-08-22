import {Request,Response,NextFunction} from "express";
import {StatusCodes} from "http-status-codes";
import {getAvailableRoomsService, updateBookingIdToRoomsService} from "../services/room.service";

export async function getAvailableRoomHandler(req: Request, res: Response, next: NextFunction) {
    const rooms = await getAvailableRoomsService(req.body);

    res.status(StatusCodes.OK).json({
        message: "Rooms found successfully",
        success: true,
        data: rooms
    })
}

export async function updateBookingIdToRoomsHandler(req: Request, res: Response, next: NextFunction) {
    const response = await updateBookingIdToRoomsService(req.body);
    res.status(StatusCodes.OK).json({
        message: "Booking ID updated to rooms successfully",
        success: true,
        data: response
    })
}