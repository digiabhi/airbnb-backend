# HotelService

A Node.js (TypeScript) microservice responsible for hotel inventory management and room availability. It exposes APIs to query available rooms for a given category and date range and to attach a booking ID to specific room-date entries. It also includes background workers and a scheduler to maintain/extend room availability windows.

## Features
- Room availability search by roomCategoryId and date range
- Update bookingId on selected room-date entries when a booking is created
- BullMQ worker to generate/extend room availability
- Scheduler to automatically extend room availability window (cron based)
- Sequelize ORM with MySQL/Postgres-compatible usage
- Structured error handling and correlation ID middleware
- Logging with winston

## Tech Stack
- Runtime: Node.js + TypeScript
- Framework: Express
- ORM: Sequelize
- Queue: BullMQ (Redis)
- Cache/Queue Backend: Redis
- Validation: zod

## Environment Variables
Create a `.env` file in `HotelService/`.

- PORT: default 3001
- REDIS_HOST: default localhost
- REDIS_PORT: default 6379
- ROOM_CRON: cron expression for scheduler (default `0 2 * * *` daily at 02:00)
- DB_HOST, DB_USER, DB_PASSWORD, DB_NAME: database connectivity

Example `.env`:
```
PORT=3000
REDIS_HOST=localhost
REDIS_PORT=6379
ROOM_CRON=0 2 * * *
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root
DB_NAME=hotel_db
```

## Running Locally
From the `HotelService` directory:

1. Install deps: `npm i`
2. Start dev server: `npm run dev`
3. On startup, the service authenticates the DB connection, sets up the room generation worker, and starts the availability extension scheduler.

## API Reference (v1)
Base: `http://localhost:<PORT>/api/v1`

- POST /rooms/available
  - Body: { roomCategoryId: number, checkInDate: string, checkOutDate: string }
  - Returns: array of available room-date records between the dates (inclusive) with bookingId = null

- POST /rooms/update-booking-id
  - Body: { bookingId: number, roomIds: number[] }
  - Action: updates bookingId on the provided room record IDs

Additional routers present:
- /ping → health/ping
- /room-generation → endpoints related to room pre-generation (if exposed)
- /scheduler → endpoints to control the scheduler (if exposed)

## Data Notes
- Room records contain at least: id, roomCategoryId, dateOfAvailability, bookingId, deletedAt
- findByRoomCategoryIdAndDateRange uses an inclusive [checkInDate, checkOutDate] filter where bookingId is null.

## Troubleshooting
- DB connection issues: verify DB_* variables and the DB is reachable
- No available rooms returned: ensure seeding or worker has generated future availability
- Redis issues for queues/scheduler: verify REDIS_HOST/REDIS_PORT
- Cron not triggering: confirm ROOM_CRON and that the scheduler is started on boot

## Scripts
- `npm run dev` — start with nodemon
- `npm start` — start with ts-node
