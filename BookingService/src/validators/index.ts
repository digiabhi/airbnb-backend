import { Request, Response, NextFunction } from 'express';
import { AnyZodObject } from 'zod';

export const validateRequestBody = (schema: AnyZodObject) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    try {
      await schema.parseAsync(req.body);
      next();
    } catch (error) {
      // If the validation fails
      res.status(400).json({
        message: 'invalid request body',
        success: false,
        error: error,
      });
    }
  };
};

export const validateQueryParams = (schema: AnyZodObject) => {
  return async (req: Request, res: Response, next: NextFunction) => {
    try {
      await schema.parseAsync(req.body);
      next();
    } catch (error) {
      // If the validation fails
      res.status(400).json({
        message: 'invalid query params',
        success: false,
        error: error,
      });
    }
  };
};
