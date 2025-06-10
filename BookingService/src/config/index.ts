// This file contains all the basic configuration logic for the app server to work.
import dotenv from 'dotenv';

dotenv.config();

type ServerConfig = {
  PORT: number;
};

export const serverConfig: ServerConfig = {
  PORT: Number(process.env.PORT) || 3001,
};
