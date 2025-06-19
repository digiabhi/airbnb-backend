import express from 'express';
import { serverConfig } from './config';
import v1Router from './routers/v1/index.router';
import v2Router from './routers/v2/index.router';
import {
  appErrorHandler,
  genericErrorHandler,
} from './middlewares/error.midleware';
import { attachCorrelationIdMiddleware } from './middlewares/correlation.middleware';
import logger from './config/logger.config';
import { setupMailerWorker } from './processors/email.processor';
import { addEmailToQueue } from './producers/email.producer';

const app = express();

app.use(express.json());
app.use(attachCorrelationIdMiddleware);
app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

app.use(appErrorHandler);
app.use(genericErrorHandler);

app.listen(serverConfig.PORT, async () => {
  logger.info(`Server is running on http://localhost:${serverConfig.PORT}`);
  setupMailerWorker();
  logger.info('Mailer worker has been set up');

  addEmailToQueue({
    to: 'singhabhishek111213@gmail.com',
    subject: 'Test Email',
    templateId: 'welcome',
    params: {
      name: 'Abhishek Singh',
      appName: 'Airbnb.com',
    },
  });
});
