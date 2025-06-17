// This file contains all the basic configuration logic for the app server to work.
import dotenv from 'dotenv';

dotenv.config();

type ServerConfig = {
  PORT: number;
  REDIS_SERVER_URL: string;
  LOCK_TTL: number;
};

export const serverConfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
  REDIS_SERVER_URL: process.env.REDIS_SERVER_URL || 'redis://localhost:6379',
  LOCK_TTL: Number(process.env.LOCK_TTL) || 5000, // Default to 5 seconds
};
