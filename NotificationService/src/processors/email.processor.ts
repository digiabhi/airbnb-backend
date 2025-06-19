import { Job, Worker } from 'bullmq';
import { MAILER_QUEUE } from '../queues/mailer.queue';
import { getRedisConnObject } from '../config/redis.config';
import { MAILER_PAYLOAD } from '../producers/email.producer';
import { NotificationDTO } from '../dto/notification.dto';
import { renderMailTemplate } from '../templates/templates.handler';
import { sendEmail } from '../services/mailer.service';
import logger from '../config/logger.config';

export const setupMailerWorker = () => {
  const emailProcessor = new Worker<NotificationDTO>(
    MAILER_QUEUE, // Name of the queue
    async (job: Job) => {
      if (job.name !== MAILER_PAYLOAD) {
        throw new Error('Invalid job name');
      }
      // Call the service layer
      const payload = job.data;

      console.log(`Processing email payload: ${JSON.stringify(payload)}`);

      const emailContent = await renderMailTemplate(
        payload.templateId,
        payload.params
      );

      await sendEmail(payload.to, payload.subject, emailContent);

      logger.info(
        `Email sent to ${payload.to} with subject "${payload.subject}"`
      );
    }, // Process function
    {
      connection: getRedisConnObject(),
    }
  );
  emailProcessor.on('error', (error) => {
    console.error('Error processing email:', error);
  });

  emailProcessor.on('failed', () => {
    console.error('Email processing failed');
  });

  emailProcessor.on('completed', () => {
    console.log('Email processing completed successfully');
  });
};
