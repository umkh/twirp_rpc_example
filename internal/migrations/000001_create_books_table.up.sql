CREATE TABLE IF NOT EXISTS book
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) UNIQUE NOT NULL,
    author     VARCHAR(255)        NOT NULL,
    created_at timestamptz default current_timestamp
);