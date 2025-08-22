# AuthService

A Go-based authentication and authorization microservice for the Airbnb-Backend monorepo. It provides user signup/login with JWT, role management, permission assignment, and route protection via middleware. Built with chi, MySQL, and bcrypt.

## Features
- User management: signup, login, profile
- JWT authentication with middleware
- Role management: CRUD roles
- Permissions: assign/remove permissions to roles, list role permissions
- Role assignment: assign roles to users
- Reverse proxy utility for upstream services with user context header
- Request validation and unified JSON response format

## Tech Stack
- Language: Go
- Web framework: github.com/go-chi/chi
- DB: MySQL (github.com/go-sql-driver/mysql)
- JWT: github.com/golang-jwt/jwt/v5
- Validation: github.com/go-playground/validator/v10
- Env: github.com/joho/godotenv
- Crypto: golang.org/x/crypto/bcrypt

## Project Structure (high-level)
- app/ — application bootstrapping (config, HTTP server)
- config/env — .env loading and typed getters
- config/db — MySQL connection setup
- controllers/ — HTTP handlers
- dto/ — request DTOs with validation tags
- middlewares/ — JWT auth, role checks, validators
- router/ — route wiring per resource (users, roles)
- db/repositories/ — repositories for users, roles, role_permissions, user_roles
- models/ — domain models
- utils/ — JSON helpers, password, proxy

## Getting Started

### Prerequisites
- Go 1.21+ (go.mod currently targets go 1.24 line; use latest stable Go)
- MySQL 8+ accessible from the service

### Environment Variables
Create an `.env` file in `AuthService/` or set env vars in your environment. A sample is provided in `.env.sample`.

Required/common variables:
- PORT — default ":3000" (example: ":8080"). Note the leading colon is expected.
- DB_USER — default "root"
- DB_PASS — default "root"
- DB_NET — default "tcp"
- DB_ADDR — default "127.0.0.1:3306"
- DB_NAME — default "auth_db"
- JWT_SECRET — secret used to sign/verify JWTs (default fallback exists but you should set this)

Example `.env`:
```
PORT=":8080"
DB_USER="root"
DB_PASS="root"
DB_NET="tcp"
DB_ADDR="127.0.0.1:3306"
DB_NAME="auth_db"
JWT_SECRET="replace-with-a-strong-secret"
```

### Database
This service expects the following tables (names inferred from queries):
- users(id, username, email, password, created_at, updated_at)
- roles(id, name, description, created_at, updated_at)
- permissions(id, name, description, resource, action, created_at, updated_at)
- role_permissions(id, role_id, permission_id, created_at, updated_at)
- user_roles(user_id, role_id)

Ensure these tables exist before running (migrations are not included in this service).

### Install and Run
From the `AuthService` directory:
- Install dependencies (Go will fetch modules on build)
- Run in dev:
```
go run ./main.go
```
- Or build and run:
```
go build -o authservice .
./authservice
```
The server will start on the configured PORT (e.g., :8080).

## API Reference
Base URL: http://localhost:8080 (adjust port if changed; omit scheme/host if running behind a gateway)

General Responses
- Success: { status: "success", message: string, data: any }
- Error: { status: "error", message: string, error: string }

### Health
- GET /ping → 200 text: "pong"

### Auth & Users
- POST /signup
  - Body: { username: string(min 3), email: valid email, password: string(min 8) }
  - 201 → success with created user (id, username, email)
- POST /login
  - Body: { email, password }
  - 200 → success with JWT token string in data
- GET /profile
  - Headers: Authorization: Bearer <JWT>
  - Access: any of roles: user, admin
  - Query: optional id=<userId> (falls back to id in JWT claims)
  - 200 → user details

JWT Claims on login
- id: number
- email: string

### Roles
- GET /roles → list roles
- GET /roles/{id} → role by id
- POST /roles
  - Body: { name: string(3-50), description: string(5-200) }
- PUT /roles/{id}
  - Body: { name, description }
- DELETE /roles/{id}

#### Role Permissions
- GET /roles/{id}/permissions → list permissions for a role
- POST /roles/{id}/permissions
  - Body: { permission_id: number }
- DELETE /roles/{id}/permissions/{permissionId}

#### Assign Role to User
- POST /roles/{userId}/assign/{roleId}
  - Headers: Authorization: Bearer <JWT>
  - Access: requires role: admin (RequireAllRoles("admin"))

### Reverse Proxy Utility
- Any request to /fakestoreservice/* is proxied to https://fakestoreapi.in with the prefix stripped.
- If authenticated, the middleware will inject header X-User-ID with the JWT user id to the upstream.

## Middleware
- JWTAuthMiddleware — parses Bearer token using JWT_SECRET, populates context with userID and email
- RequireAllRoles / RequireAnyRole — checks roles via DB queries (user_roles, roles)
- Request validators — parse and validate JSON DTOs, place payload in request context

## Running Tests
There are no unit tests provided in this module at present.

## Troubleshooting
- Cannot connect to DB: verify DB_ADDR, DB_USER, DB_PASS, DB_NAME; ensure MySQL is reachable and schema exists
- Unauthorized: ensure Authorization header is "Bearer <token>"
- 403 Forbidden on role-protected routes: ensure user has required roles in user_roles table
- Token verification error: ensure JWT_SECRET used at login is the same across services/instances

