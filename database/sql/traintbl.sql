CREATE TABLE public.trains
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