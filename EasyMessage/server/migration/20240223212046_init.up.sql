CREATE TABLE IF NOT EXISTS users(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    short_name text NOT NULL,
    email text NOT NULL UNIQUE,
    password text NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms(
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    first_user_id bigint NOT NULL,
    second_user_id bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS messages(
    id bigserial PRIMARY KEY,
    room_id bigint NOT NULL ,
    user_id bigint NOT NULL,
    content text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);