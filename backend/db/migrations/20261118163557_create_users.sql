-- +goose Up
CREATE TABLE users (
    id            uuid        PRIMARY KEY DEFAULT uuidv7(),
    email         text        NOT NULL,
    password_hash text        NOT NULL,
    role          text        NOT NULL DEFAULT 'USER' CHECK (role IN ('USER', 'ADMIN')),
    created_at    timestamptz NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX users_email_lower_key ON users (lower(email));

-- +goose Down
DROP TABLE users;
