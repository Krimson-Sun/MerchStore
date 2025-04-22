-- +goose NO TRANSACTION
-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255)
);
CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_id UUID NOT NULL,
    token TEXT NOT NULL,
    expired_at TIMESTAMPTZ DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX idx_users_email ON users (email);
CREATE INDEX CONCURRENTLY idx_session_token ON sessions (token);
-- +goose Down
DROP INDEX IF EXISTS idx_session_token;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;