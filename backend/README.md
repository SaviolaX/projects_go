# Blog API (Go)

HTTP API for a small blog: users (register, login, logout with JWT) and posts (CRUD). Built with [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/), and SQLite ([glebarez/sqlite](https://github.com/glebarez/sqlite)).

## Requirements

- Go **1.25.8** or compatible (see `go.mod`)

## Configuration

Environment variables are read from the process environment. Optionally create a `.env` file in this directory ([godotenv](https://github.com/joho/godotenv) loads it on startup).

| Variable           | Required    | Default      | Description                                      |
| ------------------ | ----------- | ------------ | ------------------------------------------------ |
| `JWT_SECRET`       | **Yes**     | —            | Secret used to sign and verify JWT access tokens |
| `DB_PATH`          | **Yes**     | —            | SQLite database file path (e.g. `blog.db`)       |
| `APP_PORT`         | Recommended | _(see note)_ | Listen address passed to Gin, e.g. `:8080`       |
| `JWT_EXPIRE_HOURS` | No          | `72`         | JWT lifetime in hours                            |

**Note:** If `APP_PORT` is unset, Gin is called with an empty address; set it explicitly (for example `:8080`) so the server listens where you expect.

## Run the server

From the `backend` directory:

```bash
go run ./cmd/server/main.go
```

Or use the Makefile:

```bash
make run
```

On start, the app connects to SQLite, runs GORM auto-migrations (`User`, `Post`, `Category`), then serves HTTP.

**Base URL:** `http://localhost<APP_PORT>` — e.g. `http://localhost:3000` if `APP_PORT=:3000`.

CORS is enabled with [gin-contrib/cors](https://github.com/gin-contrib/cors) defaults so browser clients on other origins can call the API.

---

## Authentication

1. **Register** — `POST /api/v1/auth/register` (no auth).
2. **Login** — `POST /api/v1/auth/login` (no auth). Response includes a **`token`** string.
3. **Protected routes** — send the token in the header:

   `Authorization: Bearer <token>`

4. **Logout** — `POST /api/v1/auth/logout` with the same `Authorization` header. The current token is added to an in-memory blocklist until it expires (server restart clears the blocklist).

Invalid or missing tokens on protected routes return **401** with a JSON `error` field.

---

## API endpoints

All paths are under **`/api/v1`**. JSON responses use objects with `error`, `message`, `status`, or payload keys as implemented in the handlers.

### Auth

| Method | Path                    | Auth       | Description                            |
| ------ | ----------------------- | ---------- | -------------------------------------- |
| `POST` | `/api/v1/auth/register` | No         | Create a user                          |
| `POST` | `/api/v1/auth/login`    | No         | Login; returns user fields and `token` |
| `POST` | `/api/v1/auth/logout`   | Bearer JWT | Blocklist current token                |

**Register** — JSON body:

```json
{
  "Username": "alice",
  "Email": "alice@example.com",
  "Password": "secret12"
}
```

Validation (server): username longer than 3 characters, valid email, password at least 6 characters.

**Response:** `201` with `{"status":"created"}`.

**Login** — JSON body:

```json
{
  "Username": "alice",
  "Password": "secret12"
}
```

Validation: username longer than 3 characters, password at least 6 characters.

**Response:** `200` with `id`, `username`, `email`, `createdAt`, and **`token`**.

### Posts (public)

| Method | Path                | Auth | Description                  |
| ------ | ------------------- | ---- | ---------------------------- |
| `GET`  | `/api/v1/posts`     | No   | List posts                   |
| `GET`  | `/api/v1/posts/:id` | No   | Get one post by numeric `id` |

**List response:** `200` with `{"posts":[...]}`.

**Get by ID response:** `200` with `{"post":{...}}`.

### Posts (protected)

These routes use the **`/api/v1/posts`** group with JWT middleware.

| Method   | Path                | Auth       | Description |
| -------- | ------------------- | ---------- | ----------- |
| `POST`   | `/api/v1/posts/`    | Bearer JWT | Create post |
| `PUT`    | `/api/v1/posts/:id` | Bearer JWT | Update post |
| `DELETE` | `/api/v1/posts/:id` | Bearer JWT | Delete post |

**Create** — `multipart/form-data` (not JSON) fields:

- `title` — at least 4 characters after trim
- `entry` — at least 5 characters after trim

**Response:** `201` with `{"status":"created"}`.

**Update** — `multipart/form-data` fields:

- `title`
- `entry`
- `category_id` — unsigned integer (required by handler)

**Response:** `200` with `{"status":"updated"}`.

**Delete** — no body.

**Response:** `200` with `{"status":"deleted"}`.

**Path note:** Create is registered on the group as `POST /`, which resolves to **`POST /api/v1/posts/`** (trailing slash). If a client gets 404 on `POST /api/v1/posts`, try **`POST /api/v1/posts/`**.

---

## Example: obtain a token and call a protected endpoint

```bash
export API=http://localhost:3000

# Register
curl -sS -X POST "$API/api/v1/auth/register" \
  -H 'Content-Type: application/json' \
  -d '{"Username":"alice","Email":"alice@example.com","Password":"secret12"}'

# Login and capture token (requires jq)
TOKEN=$(curl -sS -X POST "$API/api/v1/auth/login" \
  -H 'Content-Type: application/json' \
  -d '{"Username":"alice","Password":"secret12"}' | jq -r .token)

# Create post (multipart)
curl -sS -X POST "$API/api/v1/posts/" \
  -H "Authorization: Bearer $TOKEN" \
  -F 'title=My first post' \
  -F 'entry=Hello from the API.'

# List posts (no auth)
curl -sS "$API/api/v1/posts"
```

---

## Project layout

- `cmd/server/main.go` — entrypoint: config, DB, wiring, HTTP server
- `internal/handler` — HTTP routes and handlers (`router.go` defines endpoints)
- `internal/service` — business logic
- `internal/repository` — GORM data access
- `internal/model` — entities
- `internal/dto` — request/response shapes
- `internal/auth` — JWT and token blocklist
- `internal/middleware` — `AuthMiddleware`
- `internal/config` — env-based configuration
- `internal/db` — SQLite connection and migrations

Module path: `github.com/SaviolaX/blog` (see `go.mod`).
