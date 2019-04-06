CREATE TABLE public.trip
(
  trip_id uuid NOT NULL,
  trip_name character varying(30) NOT NULL,
  trip_ticket_id uuid NOT NULL,
  trip_return_ticket_id uuid,
  hotel_ticket_id uuid,
  total_trip_price double precision NOT NULL,
  CONSTRAINT trip_pkey PRIMARY KEY (trip_id)
)