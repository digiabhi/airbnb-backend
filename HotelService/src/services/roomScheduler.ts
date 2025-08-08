import * as cron from 'node-cron';
import {RoomRepository} from "../repositories/room.repository";
import RoomCategoryRepository from "../repositories/roomCategory.repository";

const roomRepository = new RoomRepository();
const roomCategoryRepository = new RoomCategoryRepository();

export class RoomSchedulerService {
    private cronJob: cron.ScheduledTask | null = null;


}