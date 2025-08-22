# NotificationService

A Node.js (TypeScript) microservice that sends email notifications using a BullMQ queue and Nodemailer. A worker processes queued email jobs and sends templated emails (Handlebars supported).

## Features
- Email queue using BullMQ (Redis backend)
- Nodemailer transport (Gmail-supported via app password)
- Handlebars templates support
- Structured error handling and correlation ID middleware
- Logging with winston

## Tech Stack
- Runtime: Node.js + TypeScript
- Framework: Express
- Queue: BullMQ (Redis)
- Mailer: Nodemailer
- Templates: Handlebars
- Validation: zod

## Environment Variables
Create a `.env` file in `NotificationService/`.

- PORT: default 3001
- REDIS_HOST: default localhost
- REDIS_PORT: default 6379
- MAIL_USER: sender email address (e.g. Gmail)
- MAIL_PASS: email app password (for Gmail, create an App Password)

Example `.env`:
```
PORT=3002
REDIS_HOST=localhost
REDIS_PORT=6379
MAIL_USER=your.email@gmail.com
MAIL_PASS=your-app-password
```

## Running Locally
From the `NotificationService` directory:

1. Install deps: `npm i`
2. Start dev server: `npm run dev`
3. On startup, the mailer worker is initialized and a sample email job is enqueued (see src/server.ts → addEmailToQueue usage).

## API and Queue Usage
- The service exposes v1 routes like `/api/v1/ping` for health checks.
- To enqueue an email programmatically from within the service or from other services:

  Example payload (NotificationDTO):
  ```ts
  addEmailToQueue({
    to: 'recipient@example.com',
    subject: 'Welcome',
    templateId: 'welcome',
    params: { name: 'Jane', appName: 'Airbnb.com' }
  });
  ```

- Worker: processes queue jobs and sends emails using the transporter configured in `src/config/mailer.config.ts`.

## Troubleshooting
- Gmail blocks sign-in: ensure you are using an App Password and 2FA is enabled on the account
- Redis not reachable: verify REDIS_HOST/REDIS_PORT and that Redis is running
- Emails not sent: check transporter config and MAIL_USER/MAIL_PASS
- Templates missing: ensure the template files exist and templateId matches

## Scripts
- `npm run dev` — start with nodemon
- `npm start` — start with ts-node
