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
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  IF NOT EXISTS trains
(
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    departure_city text NOT NULL,
    arrival_city text NOT NULL,
    departure_time text,
    departure_date text,
    arrival_time text,
    arrival_date text,
    PRIMARY KEY (id)
)
CREATE TABLE  IF NOT EXISTS planes (
    id uuid NOT NULL,
    departureCity text NOT NULL,
    arrivalCity text NOT NULL,
    PRIMARY KEY (id)
)
CREATE TABLE IF NOT EXISTS trips (
  TripID uuid NOT NULL,
  TripName character varying(30) NOT NULL,
  TripTicketID uuid NOT NULL,
  TripReturnTicketID uuid,
  TotalTripPrice double precision NOT NULL,
  CONSTRAINT trip_pkey PRIMARY KEY (TripID)
)
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

CREATE TABLE IF NOT EXISTS events (
id uuid DEFAULT uuid_generate_v1(),
title VARCHAR (64) NOT NULL,
category VARCHAR (64) NOT NULL,
town VARCHAR (64) NOT NULL,
date DATE ,
price INT ,
PRIMARY KEY (id)
);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE  IF NOT EXISTS orders (
    UserId uuid REFERENCES users(id) NOT NULL,
    TicketId uuid REFERENCES tickets(id) ,
    Amount     numeric NOT NULL DEFAULT 1,
    CONSTRAINT orders_pkey PRIMARY KEY (UserId,TicketId) );

INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Work Fair for Students','fair','Lviv','2019-04-19', 0);
 INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Malevich','festival','Lviv','2019-04-24', 700);
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'BlockChainUA','conference','Kyiv','2019-04-28', 3405);
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Film Presentation','entertaiment','Kyiv','2019-04-16', 150);
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Pinchuk Art House','entertaiment','Kyiv','2019-04-09', 350);
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Hamlet','theatre','Kyiv','2019-04-22', 100);
INSERT INTO events (id, title, category, town, date, price)
 VALUES (uuid_generate_v1(),'Scorpions','concert','Kyiv','2019-04-19', 649);

INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (23, 'museum', '-10%', 23.23, 22.99, 'Frederik', 'Lonkardi');

INSERT INTO tickets(place, ticket_type, discount, price, total_price, name,
                    surname)
VALUES (3412, 'Museum', '', 5.23, 5.23, 'Valentina', 'Kit');

INSERT INTO trips
  (TripID, TripName, TripTicketID, TripReturnTicketID, TotalTripPrice)
  VALUES ('b977aef0-64ee-11e9-a923-1681be663d3e','Relax','e4cd6b08-64ee-11e9-a923-1681be663d3e',
   'efac0872-64ee-11e9-a923-1681be663d3e', 1150);
INSERT INTO trips
  (TripID, TripName, TripTicketID, TripReturnTicketID, TotalTripPrice)
  VALUES ('15961f82-64ef-11e9-a923-1681be663d3e','Chill','25d5f700-64ef-11e9-a923-1681be663d3e',
  '2b30682a-64ef-11e9-a923-1681be663d3e', 950);
INSERT INTO trips
  (TripID, TripName, TripTicketID, TripReturnTicketID, TotalTripPrice)
  VALUES ('3c180ee0-64ef-11e9-a923-1681be663d3e','CoolWeek','434a5614-64ef-11e9-a923-1681be663d3e',
  '56c90c3a-64ef-11e9-a923-1681be663d3e', 850);

INSERT INTO planes
  (id,  departureCity,arrivalCity,)
  VALUES ('db593603-f349-4489-9e1c-0fbeef1dd4f7','Lviv','Kyiv');
INSERT INTO planes
  (id,  departureCity,arrivalCity,)
  VALUES ('7a723e05-4187-45b4-a598-7e14e4940d99','Kyiv','Kharkiv');
INSERT INTO planes
  (id,  departureCity,arrivalCity,)
  VALUES ('b6a6e279-9a4b-4d6c-9a4d-3d8c1e3b3658','Ivano-Frankivsk','Odessa');

INSERT INTO trains
  (id,  departure_city, arrival_city, departure_time, departure_date,arrival_time,arrival_date)
  VALUES ( uuid_generate_v4(), 'Lviv','Kyiv' ,'22:30:00','2019-04-26', '05:55:00','2019-04-27');

  INSERT INTO trains
  (id,  departure_city, arrival_city, departure_time, departure_date,arrival_time,arrival_date)
  VALUES ( uuid_generate_v4(), 'Kyiv','Harkiv' ,'18:30:00','2019-04-27', '23:55:00','2019-04-27');

  INSERT INTO trains
  (id,  departure_city, arrival_city, departure_time, departure_date,arrival_time,arrival_date)
  VALUES ( uuid_generate_v4(), 'Uzgorod','Lviv' ,'07:30:00','2019-04-28', '16:55:00','2019-04-28');


INSERT INTO users
 (id, name, surname, login, password, role, email)
VALUES ('207a1a86-5d89-41d0-aa7e-c589cdc2a39e', 'Oksana', 'Zhykina','oks_zh','$2a$14$3JFqIzSAXhHk8Opq0/BSxuSWkeiZCiLo2gXmeKt0pQ11MU4YY8O
        /K','admin','oks88zh@gmail.com');

INSERT INTO users
 (id, name, surname, login, password, role, email)
VALUES ('207a1a86-5d89-41d0-aa7e-c589cdc2a39e', 'Oksana', 'Zhykina','oks_zh','$2a$14$3JFqIzSAXhHk8Opq0/BSxuSWkeiZCiLo2gXmeKt0pQ11MU4YY8O
        /K','admin','oks88zh@gmail.com');


INSERT INTO orders
  (UserID, TicketID)
  VALUES ('207a1a86-5d89-41d0-aa7e-c589cdc2a39e','e1e0fd8d-9645-4c46-8ec7-8d27025a5ee8');

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table IF EXISTS users,trains,planes,trips,tickets,events,orders;