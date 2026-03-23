# Backend Quick Start

## 1) Environment

Copy env template:

```bash
cp .env.example .env
```

You can use either local or remote PostgreSQL:

- Local: set `DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME/DB_SSLMODE`
- Remote: set `DATABASE_URL` directly

Default app port is `8979`.

## 2) Initialize database

Run SQL script:

```bash
psql "$DATABASE_URL" -f migrations/0001_init.sql
psql "$DATABASE_URL" -f migrations/0002_site_settings.sql
```

Or use host/user style login and run the same SQL file.

Initial admin account inserted by SQL:

- username: `admin`
- password: `myhx0308` (stored as bcrypt hash)

## 3) Run server

```bash
go mod tidy
go run ./cmd/server
```

Health check:

```bash
GET http://localhost:8979/healthz
```

## 4) Implemented endpoints

- `POST /api/v1/auth/login`
- `GET /api/v1/auth/me`
- `POST /api/v1/auth/logout`
- `GET /api/v1/posts`
- `GET /api/v1/posts/:id`
- `GET /api/v1/posts/slug/:slug`
- `GET /api/v1/tags`
- `GET /api/v1/categories`
- `GET /api/v1/timeline`
- `GET /api/v1/site/settings`
- `GET /api/v1/admin/posts`
- `GET /api/v1/admin/posts/:id`
- `GET /api/v1/admin/posts/check-slug`
- `POST /api/v1/admin/posts`
- `PUT /api/v1/admin/posts/:id`
- `DELETE /api/v1/admin/posts/:id`
- `PATCH /api/v1/admin/posts/:id/publish`
- `PATCH /api/v1/admin/posts/:id/unpublish`
- `PATCH /api/v1/admin/posts/:id/top`
- `GET /api/v1/admin/tags`
- `POST /api/v1/admin/tags`
- `PUT /api/v1/admin/tags/:id`
- `DELETE /api/v1/admin/tags/:id`
- `GET /api/v1/admin/categories`
- `POST /api/v1/admin/categories`
- `PUT /api/v1/admin/categories/:id`
- `DELETE /api/v1/admin/categories/:id`
- `POST /api/v1/admin/upload/image`
- `POST /api/v1/admin/upload/images`
- `GET /api/v1/admin/media`
- `DELETE /api/v1/admin/media/:id`
- `GET /api/v1/admin/site/settings`
- `PUT /api/v1/admin/site/settings`
- `GET /api/v1/admin/dashboard`

## 5) Swagger generation

After installing `swag`:

```bash
swag init -g cmd/server/main.go -o docs
```

Swagger UI:

```text
http://localhost:8979/swagger/index.html
```
