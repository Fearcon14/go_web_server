-- +goose Up
-- This section contains SQL that will be executed when running the migration UP (applying the migration)
-- This is where we create new tables, add columns, etc.

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL UNIQUE
);

-- +goose Down
-- This section contains SQL that will be executed when running the migration DOWN (rolling back the migration)
-- This should reverse whatever was done in the Up section

DROP TABLE users;

