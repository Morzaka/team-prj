-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- Tickets
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  IF NOT EXISTS tickets
(
    id             uuid        DEFAULT uuid_generate_v4() NOT NULL,
    train_id       uuid        DEFAULT uuid_generate_v4() NOT NULL, -- references train(id), -- uncomment when train,
    plane_id       uuid        DEFAULT uuid_generate_v4() NOT NULL, -- references plane(id), -- plane and users tables
    user_id        uuid        DEFAULT uuid_generate_v4() NOT NULL, -- references user(id),  -- will be exist.
    place          INTEGER                                NOT NULL,
    ticket_type    VARCHAR(30) DEFAULT 'Train'            NOT NULL,
    discount       VARCHAR(10) DEFAULT '-0%'              NOT NULL,
    price          DECIMAL(5, 2)                          NOT NULL,
    total_price    DECIMAL(5, 2)                          NOT NULL,
    name           VARCHAR(30),
    surname        VARCHAR(30),
    from_place     VARCHAR(30),
    departure_date DATE,
    departure_time TIME,
    to_place       VARCHAR(30),
    arrival_date   DATE,
    arrival_time   TIME,
    PRIMARY KEY (id)
);

INSERT INTO tickets(id,train_id,plane_id,user_id,place,ticket_type,discount, price, total_price, "name",surname)
VALUES ('e1e0fd8d-9645-4c46-8ec7-8d27025a5ee8','84458656-8f8c-40a7-827a-7bc14cf86314',
'98a57b8e-c081-4234-9c29-f29e74f82221','c1fcc072-b47d-43de-94a7-26208064ce1a',
44, 'bus', '-5%', 200.00, 180.00, 'Stepko', 'Brovarskij');


INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (23, 'museum', '-10%', 23.23, 22.99, 'Frederik', 'Lonkardi');

INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (3412, 'Museum', '', 5.23, 5.23, 'Valentina', 'Kit');
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table tickets