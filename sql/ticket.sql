-- Tickets

CREATE TABLE tickets
(
    id             uuid DEFAULT uuid_generate_v4() NOT NULL,
    train_id       uuid DEFAULT uuid_generate_v4() NOT NULL, -- references train(id),
    plane_id       uuid DEFAULT uuid_generate_v4() NOT NULL, -- references plane(id),
    users_id       uuid DEFAULT uuid_generate_v4() NOT NULL, -- references users(id),
    place          SMALLINT NOT NULL,
    type           VARCHAR(30) DEFAULT 'Train' NOT NULL,
    discount       VARCHAR(10) DEFAULT '-0%' NOT NULL,
    price          VARCHAR(10) NOT NULL,
    total_price    VARCHAR(10) NOT NULL,
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

INSERT INTO tickets(
    place, type, discount, price, total_price, name, surname, from_place, departure_date, departure_time, to_place, arrival_date, arrival_time)
VALUES (23, 'Plane', '-20%', '100', '80', 'Lyubomyr', 'Mykhalchyshyn', 'Lviv', '04-apr-2019', '22:30', 'Kharkiv', '05-apr-2019', '07:30');

INSERT INTO tickets(
    place, discount, price, total_price, name, surname, from_place, departure_date, departure_time, to_place, arrival_date, arrival_time)
VALUES (44, '-5%', '200', '180', 'Stepko', 'Brovarskij', 'Zhmerenka', '12-may-2019', '17:23', 'Lohinka', '12-may-2019', '23:56');

INSERT INTO tickets(
    place, type, price, total_price, name, surname)
VALUES (3412, 'Museum',  '5', '5', 'Valentina', 'Kit');