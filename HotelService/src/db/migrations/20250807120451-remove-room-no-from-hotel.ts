'use strict';

import {QueryInterface} from "sequelize";

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface: QueryInterface) {
      await queryInterface.sequelize.query(
          `ALTER TABLE rooms DROP COLUMN room_no;`
      )
  },

  async down (queryInterface: QueryInterface) {
      await queryInterface.sequelize.query(
          `ALTER TABLE rooms ADD COLUMN room_no INT NOT NULL ;`
      )
  }
};
