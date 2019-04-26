-- +goose Up
-- SQL in this section is executed when the migration is applied.
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

INSERT INTO trains
  (id,  departure_city, arrival_city, departure_time, departure_date,arrival_time,arrival_date)
  VALUES ( uuid_generate_v4(), 'Lviv','Kyiv' ,'22:30:00','2019-04-26', '05:55:00','2019-04-27');

  INSERT INTO trains
  (id,  departure_city, arrival_city, departure_time, departure_date,arrival_time,arrival_date)
  VALUES ( uuid_generate_v4(), 'Kyiv','Harkiv' ,'18:30:00','2019-04-27', '23:55:00','2019-04-27');

  INSERT INTO trains
  (id,  departure_city, arrival_city, departure_time, departure_date,arrival_time,arrival_date)
  VALUES ( uuid_generate_v4(), 'Uzgorod','Lviv' ,'07:30:00','2019-04-28', '16:55:00','2019-04-28');
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

drop table trains;