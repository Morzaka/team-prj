CREATE TABLE public.trip
(
  TripID uuid NOT NULL,
  TripName character varying(30) NOT NULL,
  TripTicketId uuid NOT NULL,
  TripReturnTicketId uuid,
  TotalTripPrice double precision NOT NULL,
  CONSTRAINT trip_pkey PRIMARY KEY (TripID)
)