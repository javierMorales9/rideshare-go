-- +migrate Up
CREATE TABLE locations (
    id         SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    address    TEXT         NOT NULL UNIQUE,
    state      CHAR(2)      NOT NULL,
    latitude   NUMERIC(9,6),
    longitude  NUMERIC(9,6),

    CHECK (char_length(state) = 2)
);

-- Búsqueda rápida por dirección (ILIKE '%foo%')
--CREATE INDEX idx_locations_address_trgm ON locations USING gin (address gin_trgm_ops);
