-- +migrate Up
CREATE TABLE users (
    id                      SERIAL PRIMARY KEY,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at              TIMESTAMPTZ,

    first_name              VARCHAR(100) NOT NULL,
    last_name               VARCHAR(100) NOT NULL,
    email                   VARCHAR(255) NOT NULL UNIQUE,
    password_hash           VARCHAR(60)  NOT NULL,
    type                    VARCHAR(20)  NOT NULL,
    drivers_license_number  VARCHAR(100),

    CHECK (type IN ('driver', 'rider'))
);

CREATE INDEX idx_users_type ON users(type);
