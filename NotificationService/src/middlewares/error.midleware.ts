import { Request, Response, NextFunction } from 'express';
import { AppError } from '../utils/errors/app.error';

export const appErrorHandler = (
  err: AppError,
  req: Request,
  res: Response,
  next: NextFunction
) => {
  res.status(err.statusCode).json({
    success: false,
    message: err.message,
  });
};

export const genericErrorHandler = (
  err: Error,
  req: Request,
  res: Response,
  next: NextFunction
) => {
  res.status(500).json({
    success: false,
    message: 'Internal Server Error',
  });
};
