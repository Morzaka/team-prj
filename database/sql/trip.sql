CREATE TABLE public.trip
(
  TripID uuid NOT NULL,
  TripName character varying(30) NOT NULL,
  TripTicketID uuid NOT NULL,
  TripReturnTicketID uuid,
  TotalTripPrice double precision NOT NULL,
  CONSTRAINT trip_pkey PRIMARY KEY (TripID)
)