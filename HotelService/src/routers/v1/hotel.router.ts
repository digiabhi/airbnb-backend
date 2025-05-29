import express from 'express';
import {
  createHotelHandler,
  deleteHotelHandler,
  getAllHotelsHandler,
  getHotelByIdHandler,
} from '../../controllers/hotel.controller';
import { validateRequestBody } from '../../validators';
import { hotelSchema } from '../../validators/hotel.validator';

const hotelRouter = express.Router();

hotelRouter.get('/', getAllHotelsHandler);

hotelRouter.post('/', validateRequestBody(hotelSchema), createHotelHandler);

hotelRouter.get('/:id', getHotelByIdHandler);

hotelRouter.delete('/:id', deleteHotelHandler);

export default hotelRouter;
