-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  IF NOT EXISTS users(
    id uuid NOT NULL,
    name text NOT NULL,
    surname text,
    login text NOT NULL,
    password text NOT NULL,
    role text NOT NULL DEFAULT User,
    email text NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id)
)


INSERT INTO users
 (id, name, surname, login, password, role, email)
VALUES ('207a1a86-5d89-41d0-aa7e-c589cdc2a39e', 'Oksana', 'Zhykina','oks_zh','$2a$14$3JFqIzSAXhHk8Opq0/BSxuSWkeiZCiLo2gXmeKt0pQ11MU4YY8O
        /K','admin','oks88zh@gmail.com');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table users;