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
import sequelize from './db/models/sequelize';
import {setupRoomGenerationWorker} from "./processors/roomGeneration.processor";
import { startScheduler } from './scheduler/roomScheduler';


const app = express();

app.use(express.json());
app.use(attachCorrelationIdMiddleware);
app.use('/api/v1', v1Router);
app.use('/api/v2', v2Router);

app.use(appErrorHandler);
app.use(genericErrorHandler);

app.listen(serverConfig.PORT, async () => {
  logger.info(`Server is running on http://localhost:${serverConfig.PORT}`);
  await sequelize.authenticate(); // Test the connection to the DB.
  logger.info('Database connected successfully!');
  setupRoomGenerationWorker();
  startScheduler();
  logger.info('Room availability extension scheduler initialized');
});
