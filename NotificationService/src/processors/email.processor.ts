import { Job, Worker } from 'bullmq';
import { MAILER_QUEUE } from '../queues/mailer.queue';
import { getRedisConnObject } from '../config/redis.config';
import { MAILER_PAYLOAD } from '../producers/email.producer';
import { NotificationDTO } from '../dto/notification.dto';

export const setupMailerWorker = () => {
  const emailProcessor = new Worker<NotificationDTO>(
    MAILER_QUEUE, // Name of the queue
    async (job: Job) => {
      if (job.name !== MAILER_PAYLOAD) {
        throw new Error('Invalid job name');
      }
      // Call the service layer
    }, // Process function
    {
      connection: getRedisConnObject(),
    }
  );

  emailProcessor.on('failed', () => {
    console.error('Email processing failed');
  });

  emailProcessor.on('completed', () => {
    console.log('Email processing completed successfully');
  });
};
