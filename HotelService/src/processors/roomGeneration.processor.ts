import { Job, Worker } from 'bullmq';
import { getRedisConnObject } from '../config/redis.config';
import {ROOM_GENERATION_QUEUE} from "../queues/roomGeneration.queue";
import {ROOM_GENERATION_PAYLOAD} from "../producers/roomGeneration.producer";
import logger from '../config/logger.config';
import {RoomGenerationJob} from "../dto/roomGeneration.dto";
import {generateRooms} from "../services/roomGeneration.service";

export const setupRoomGenerationWorker = () => {
    const roomGenerationProcessor = new Worker<RoomGenerationJob>(
        ROOM_GENERATION_QUEUE, // Name of the queue
        async (job: Job) => {
            if (job.name !== ROOM_GENERATION_PAYLOAD) {
                throw new Error('Invalid job name');
            }
            // Call the service layer
            const payload = job.data;

            console.log(`Processing room generation payload: ${JSON.stringify(payload)}`);

            await generateRooms(payload)
            logger.info(`Room generation completed for ${JSON.stringify(payload)}`);

        }, // Process function
        {
            connection: getRedisConnObject(),
        }
    );
    roomGenerationProcessor.on('error', (error) => {
        console.error('Error processing room generation:', error);
    });

    roomGenerationProcessor.on('failed', (_, error, __) => {
        console.error('Room generation processing failed' ,error);
    });

    roomGenerationProcessor.on('completed', ()  => {
        console.log('Room generation processing completed successfully');
    });
};
