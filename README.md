# Rideshare-go

> Mini-replica of a ridesharing app (like Uber/Lyft) ported from Ruby on Rails to Go.

This project originates from **[andyatkinson/rideshare](https://github.com/andyatkinson/rideshare)** (Rails). It ports the models, services, and API logic to **Go + Gin + Gorm**, keeping the original structure and core behaviors.

---

## 🧱 Tech Stack

* **Go** 1.22+
* **Gin** (HTTP framework)
* **Gorm** (ORM) + **PostgreSQL**
* **golang-migrate** (pure SQL migrations)
* **bcrypt** (password hashing)
* **JWT** (bearer token authentication)
* **godotenv** *(optional for dev)*

---

## 📁 Repository Structure

```
rideshare-go/
├── cmd/
│   └── api/                   # server entry point and routes
├── internal/
│   ├── config/                 # environment variable loading
│   ├── db/                     # Gorm connection / helpers
│   ├── domain/
│   │   └── model/              # Gorm structs (User, Trip, etc.)
│   ├── http/
│   │   ├── handler/            # Gin handlers (auth, trips, trip_requests)
│   │   └── middleware/         # JWT auth
│   └── service/                # domain logic (TripCreator, BookReservation, etc.)
├── migrations/                 # SQL for golang-migrate
└── README.md
```

---

## ✅ Implemented Features

### Authentication

* User registration (`/auth/register`) — roles: `driver` or `rider`.
* Login (`/auth/login`) → JWT (24h expiry).
* `AuthRequired` middleware for protected routes.

### Models (Rails parity)

* **User** (role via `type`), `DisplayName()`, bcrypt.
* **Location** (`address`, `state`, `latitude`, `longitude`). *(Geocoding pending)*
* **TripRequest** (rider, start\_location, end\_location).
* **Trip** (driver, rating 1–5, `completed_at`). Validation: no rating unless completed.
* **TripPosition** (trip location tracking; no endpoints yet).
* **Vehicle** (`name` unique, `status` enum: `draft`/`published`).
* **VehicleReservation** (vehicle booking for a `TripRequest`).

### Services

* **TripCreator**: selects a driver (random) and creates a `Trip` for a `TripRequest` **within the same transaction**.
* **BookReservation**: creates a `TripRequest` and `VehicleReservation` in one transaction.

### Endpoints

* `POST  /auth/register`  → Register user.
* `POST  /auth/login`     → Get JWT.
* `GET   /api/me`         → Protected; returns `user_id` from JWT.

**TripRequests**

* `POST  /api/trip_requests` *(protected)* → Create `TripRequest` + run `TripCreator`.
* `GET   /api/trip_requests/:id` *(public or protected)* → `{ trip_request_id, trip_id }`.

**Trips**

* `GET   /api/trips` *(protected)* → List with filters (`start_location`, `driver_name`, `rider_name`).
* `GET   /api/trips/:id` *(protected)* → Trip with 60s cache.
* `GET   /api/trips/:id/details` *(protected)* → `include=driver`, `fields[driver]=...` (JSON\:API-like).
* `GET   /api/trips/my?rider_id=X` *(protected)* → Completed trips for a rider.

Rails original did not expose endpoints for **Vehicles**/**VehicleReservations** — same here.

---

## ⚙️ Setup & Run

### 1) Requirements (Ubuntu)

```bash
sudo apt update && sudo apt install -y postgresql postgresql-contrib curl jq
sudo snap install go --classic
# golang-migrate CLI
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 2) `.env`

```dotenv
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=rideshare
DB_SSLMODE=disable

DSN=host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=${DB_SSLMODE}
PORT=8080
JWT_SECRET=super-secret-change-me
```

### 3) Database & migrations

```bash
sudo -u postgres psql -c "CREATE DATABASE rideshare;"
psql -d rideshare -c "CREATE EXTENSION IF NOT EXISTS pg_trgm;"
migrate -path migrations -database "$DSN" up
```

### 4) Run API

```bash
go run ./cmd/api
```

---

## 🟨 Pending for Rails parity

1. Full-text search (`pg_search_scope` equivalent).
2. Geocoding in `Location`.
3. Counter cache (driver's trips).
4. Advanced validations (email format, case-insensitive uniqueness).
5. Strict JSON\:API serialization.

---

## 🚀 Future Development Ideas

* WebSocket/SSE for real-time `TripPosition` tracking.
* PostGIS adoption for geospatial queries.
* Overbooking prevention with exclusion constraints.
* Endpoints for Vehicles and reservations.
* Pagination, sorting, and filtering improvements.
* OpenAPI/Swagger documentation.
* Observability (structured logs, Prometheus metrics, tracing).
* Security enhancements (JWT rotation, rate limiting).
* Docker Compose & CI/CD integration.

---

## 🙌 Credits

* Original project: **andyatkinson/rideshare** (Rails).
* Go port: This implementation (Gin/Gorm/golang-migrate) with focus on functional parity and extensibility.

---

## 📄 License

MIT (same as original, if applicable).

