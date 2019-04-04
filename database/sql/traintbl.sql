CREATE TABLE public.trains
(
    id serial NOT NULL,
    departure_city text NOT NULL,
    arrival_city text NOT NULL,
    departure_time time without time zone,
    departure_date date,
    arrival_time time without time zone,
    arrival_date date,
    PRIMARY KEY (id)
)