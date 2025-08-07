-- +migrate Up
CREATE TABLE trip_requests (
    id                 SERIAL PRIMARY KEY,
    created_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    rider_id           INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    start_location_id  INTEGER NOT NULL REFERENCES locations(id),
    end_location_id    INTEGER NOT NULL REFERENCES locations(id)
    -- Vehículos y otros campos irán después
);

CREATE INDEX idx_trip_requests_rider      ON trip_requests(rider_id);
CREATE INDEX idx_trip_requests_start_loc  ON trip_requests(start_location_id);
CREATE INDEX idx_trip_requests_end_loc    ON trip_requests(end_location_id);

CREATE TABLE trips (
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    trip_request_id   INTEGER UNIQUE NOT NULL REFERENCES trip_requests(id) ON DELETE CASCADE,
    driver_id         INTEGER NOT NULL REFERENCES users(id),
    completed_at      TIMESTAMPTZ,
    rating            INTEGER CHECK (rating BETWEEN 1 AND 5)
);

CREATE INDEX idx_trips_driver      ON trips(driver_id);
CREATE INDEX idx_trips_completed   ON trips(completed_at);

CREATE TABLE trip_positions (
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    trip_id    INTEGER NOT NULL REFERENCES trips(id) ON DELETE CASCADE,
    latitude   NUMERIC(9,6) NOT NULL,
    longitude  NUMERIC(9,6) NOT NULL
);

CREATE INDEX idx_trip_positions_trip ON trip_positions(trip_id);
