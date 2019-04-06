CREATE TABLE plane
(
    id uuid NOT NULL,
    departureCity text NOT NULL,
    arrivalCity text NOT NULL,
    PRIMARY KEY (id)
)
INSERT INTO plane
  (id,  departureCity,arrivalCity,)
  VALUES ('db593603-f349-4489-9e1c-0fbeef1dd4f7','Lviv','Kyiv');
INSERT INTO plane
  (id,  departureCity,arrivalCity,)
  VALUES ('7a723e05-4187-45b4-a598-7e14e4940d99','Kyiv','Kharkiv');
INSERT INTO plane
  (id,  departureCity,arrivalCity,)
  VALUES ('b6a6e279-9a4b-4d6c-9a4d-3d8c1e3b3658','Ivano-Frankivsk','Odessa');