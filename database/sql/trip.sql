CREATE TABLE public.trip
(
  TripID uuid NOT NULL,
  TripName character varying(30) NOT NULL,
  TripTicketID uuid NOT NULL,
  TripReturnTicketID uuid,
  TotalTripPrice double precision NOT NULL,
  CONSTRAINT trip_pkey PRIMARY KEY (TripID)
)

INSERT INTO trip
  (TripID, TripName, TripTicketID, TripReturnTicketID, TotalTripPrice)
  VALUES ('b977aef0-64ee-11e9-a923-1681be663d3e','Relax','e4cd6b08-64ee-11e9-a923-1681be663d3e',
   'efac0872-64ee-11e9-a923-1681be663d3e', 1150);
INSERT INTO trip
  (TripID, TripName, TripTicketID, TripReturnTicketID, TotalTripPrice)
  VALUES ('15961f82-64ef-11e9-a923-1681be663d3e','Chill','25d5f700-64ef-11e9-a923-1681be663d3e',
  '2b30682a-64ef-11e9-a923-1681be663d3e', 950);
INSERT INTO trip
  (TripID, TripName, TripTicketID, TripReturnTicketID, TotalTripPrice)
  VALUES ('3c180ee0-64ef-11e9-a923-1681be663d3e','CoolWeek','434a5614-64ef-11e9-a923-1681be663d3e',
  '56c90c3a-64ef-11e9-a923-1681be663d3e', 850);