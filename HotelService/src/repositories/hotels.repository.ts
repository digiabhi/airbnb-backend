import logger from '../config/logger.config';
import Hotel from '../db/models/hotel';
import { NotFoundError } from '../utils/errors/app.error';
import BaseRepository from "./base.repository";

export class HotelRepository extends BaseRepository<Hotel>{
  constructor() {
    super(Hotel);
  }

  async findAll(){
    const hotels = await Hotel.findAll({
      where: {
        deletedAt: null,
      },
    });

    if (!hotels) {
      logger.error(`No Hotels Found`);
      throw new NotFoundError(`No hotels found`);
    }
    logger.info(`Hotels found: ${hotels.length}`);
    return hotels;
  }

  async softDelete(id: number) {
    const hotel = await Hotel.findByPk(id);

    if (!hotel) {
      logger.error(`Hotel not found: ${id}`);
      throw new NotFoundError(`Hotel with id ${id} not found`);
    }
    hotel.deletedAt = new Date();
    await hotel.save();
    logger.info(`Hotel soft deleted: ${hotel.id}`);
    return true;
  }
}