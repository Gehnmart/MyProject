CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    short_name text NOT NULL,
    email text NOT NULL UNIQUE,
    password text NOT NULL
);