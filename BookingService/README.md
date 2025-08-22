# BookingService

A Node.js (TypeScript) microservice responsible for creating and confirming hotel room bookings. It coordinates with HotelService to reserve available rooms and uses Redis-based distributed locks (Redlock) to prevent overbooking.

## Features
- Create booking with idempotency support
- Distributed locking with Redlock to avoid concurrent double bookings
- Fetch available rooms from HotelService and attach booking IDs back to rooms
- Booking confirmation using idempotency keys
- Structured error handling and correlation ID middleware
- Logging with winston

## Tech Stack
- Runtime: Node.js + TypeScript
- Framework: Express
- Data: Prisma Client (DB configured in prisma/schema.prisma)
- Caching/Locks: Redis + redlock
- Queue: BullMQ (future extensibility)
- HTTP: axios
- Validation: zod

## Environment Variables
Create a `.env` file in `BookingService/`.

- PORT: default 3001
- REDIS_SERVER_URL: e.g., redis://localhost:6379
- LOCK_TTL: redlock TTL in ms (default 5000)
- HOTEL_SERVICE_URL: base URL for HotelService (default http://localhost:3000/api/v1)

Example `.env`:
```
PORT=3001
REDIS_SERVER_URL=redis://localhost:6379
LOCK_TTL=5000
HOTEL_SERVICE_URL=http://localhost:3000/api/v1
```

## Running Locally
From the `BookingService` directory:

1. Install deps: `npm i`
2. Start dev server: `npm run dev`
3. The API will be available on `http://localhost:<PORT>/api` (v1 and v2 routers are mounted).

## API Overview
Note: Only key flows are highlighted based on current code.

- POST /api/v1/bookings (conceptual)
  - Calls createBookingService(createBookingDTO)
  - Flow in service:
    - Calls HotelService: POST /rooms/available with { roomCategoryId, checkInDate, checkOutDate }
    - Validates there are enough available rooms for all nights
    - Acquires Redis locks for each room resource: room:{id}
    - Creates booking in DB
    - Generates and persists idempotency key
    - Calls HotelService: POST /rooms/update-booking-id to attach bookingId to the selected room dates
    - Releases locks
  - Returns: { bookingId, idempotencyKey }

- POST /api/v1/bookings/confirm (conceptual)
  - Body: { idempotencyKey }
  - confirmBookingService() finalizes the booking associated with the idempotency key.

Your project may have specific routers/controllers for these endpoints in v1/v2; wire accordingly.

## Related Services
- HotelService: Provides room availability and room booking ID updates

## Troubleshooting
- Redis connection issues: verify REDIS_SERVER_URL and that Redis is running
- Not enough rooms available: ensure HotelService seed data has available rooms for the date range
- Lock acquisition failure: check Redis, network stability, and adjust LOCK_TTL
- HOTEL_SERVICE_URL incorrect: results in axios errors when contacting HotelService

## Scripts
- `npm run dev` — start with nodemon
- `npm start` — start with ts-node
