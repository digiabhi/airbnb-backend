import { Request, Response, NextFunction } from 'express';
import {
  createHotelService,
  deleteHotelService,
  getAllHotelsService,
  getHotelByIdService,
} from '../services/hotel.service';
import { StatusCodes } from 'http-status-codes';

export async function createHotelHandler(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await createHotelService(req.body);

  res.status(StatusCodes.CREATED).json({
    message: 'Hotel created successfully',
    data: hotelResponse,
    success: true,
  });
}

export async function getHotelByIdHandler(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await getHotelByIdService(Number(req.params.id));

  res.status(StatusCodes.OK).json({
    message: 'Hotel found successfully',
    data: hotelResponse,
    success: true,
  });
}

export async function getAllHotelsHandler(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await getAllHotelsService();

  res.status(StatusCodes.OK).json({
    message: 'Hotels found successfully',
    data: hotelResponse,
    success: true,
  });
}

export async function deleteHotelHandler(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const hotelResponse = await deleteHotelService(Number(req.params.id));

  res.status(StatusCodes.OK).json({
    message: 'Hotel deleted successfully',
    data: hotelResponse,
    success: true,
  });
}
