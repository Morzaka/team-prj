-- Tickets
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE tickets
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

-- INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
--                     surname, from_place, departure_date, departure_time,
--                     to_place, arrival_date, arrival_time)
-- VALUES (23, 'Plane', '-20%', 100.00, 80.00, 'Lyubomyr', 'Mykhalchyshyn', 'Lviv',
--         '04-apr-2019', '22:30', 'Kharkiv', '05-apr-2019', '07:30');

INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (44, 'bus', '-5%', 200.00, 180.00, 'Stepko', 'Brovarskij');

INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (23, 'museum', '-10%', 23.23, 22.99, 'Frederik', 'Lonkardi');

INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (3412, 'Museum', '', 5.23, 5.23, 'Valentina', 'Kit');