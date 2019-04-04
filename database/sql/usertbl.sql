CREATE TABLE user
(
    id uuid NOT NULL,
    name text NOT NULL,
    surname text,
    login text NOT NULL,
    password text NOT NULL,
    role text NOT NULL DEFAULT User,
    CONSTRAINT user_pkey PRIMARY KEY (id)
)
