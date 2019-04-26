-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  IF NOT EXISTS orders (
    UserId uuid REFERENCES users(id) NOT NULL,
    TicketId uuid REFERENCES tickets(id) ,
    Amount     numeric NOT NULL DEFAULT 1,
    CONSTRAINT orders_pkey PRIMARY KEY (UserId,TicketId) );

INSERT INTO orders
  (UserID, TicketID)
  VALUES ('207a1a86-5d89-41d0-aa7e-c589cdc2a39e','e1e0fd8d-9645-4c46-8ec7-8d27025a5ee8');
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table orders;