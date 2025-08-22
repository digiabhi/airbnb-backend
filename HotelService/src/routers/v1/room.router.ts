import express from 'express';
import {getAvailableRoomHandler, updateBookingIdToRoomsHandler} from "../../controllers/room.controller";
import {validateRequestBody} from "../../validators";
import {getAvailableRoomsSchema, updateBookingIdToRoomsSchema} from "../../validators/room.validator";

const roomRouter = express.Router();

roomRouter.post('/available', validateRequestBody(getAvailableRoomsSchema), getAvailableRoomHandler);
roomRouter.post('/update-booking-id', validateRequestBody(updateBookingIdToRoomsSchema), updateBookingIdToRoomsHandler);

export default roomRouter;
