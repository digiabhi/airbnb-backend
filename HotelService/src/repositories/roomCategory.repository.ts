import RoomCategory from "../db/models/roomCategory";
import BaseRepository from "./base.repository";
import {NotFoundError} from "../utils/errors/app.error";

export class RoomCategoryRepository extends BaseRepository<RoomCategory> {
    constructor() {
        super(RoomCategory);
    }

    async findAllByHotelId(hotelId:number) {
        const roomCategories = await RoomCategory.findAll({
            where: {
                hotelId: hotelId,
                deletedAt: null,
            },
        });

        if (!roomCategories || roomCategories.length === 0) {
            throw new NotFoundError(`No room categories found for hotel with id ${hotelId}`);
        }

        return roomCategories;
    }
}