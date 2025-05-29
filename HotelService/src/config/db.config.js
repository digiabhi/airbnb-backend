const dotenv = require('dotenv');

dotenv.config();

const config = {
  development: {
    username: process.env.DB_USER || 'root',
    password: process.env.DB_PASSWORD || 'root',
    database: process.env.DB_NAME || 'airbnb_dev',
    host: process.env.DB_HOST || 'localhost',
    dialect: 'mysql',
  },
};

module.exports = config;
