-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    user_name TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    currency NUMERIC NOT NULL
);

-- +goose Down
DROP TABLE users;