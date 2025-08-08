-- +migrate Up
CREATE TABLE vehicles (
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    name        VARCHAR(255) NOT NULL UNIQUE,
    status      VARCHAR(20)  NOT NULL,

    CHECK (status IN ('draft','published'))
);

CREATE INDEX idx_vehicles_status ON vehicles(status);

CREATE TABLE vehicle_reservations (
    id               SERIAL PRIMARY KEY,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    vehicle_id       INTEGER NOT NULL REFERENCES vehicles(id)        ON DELETE CASCADE,
    trip_request_id  INTEGER NOT NULL REFERENCES trip_requests(id)   ON DELETE CASCADE,

    starts_at        TIMESTAMPTZ NOT NULL,
    ends_at          TIMESTAMPTZ NOT NULL,

    CHECK (ends_at > starts_at)
);

CREATE INDEX idx_vehicle_reservations_vehicle ON vehicle_reservations(vehicle_id);
CREATE INDEX idx_vehicle_reservations_request ON vehicle_reservations(trip_request_id);
