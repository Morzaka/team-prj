CREATE TABLE IF NOT EXISTS events (
id uuid DEFAULT uuid_generate_v1(),
title VARCHAR (64) NOT NULL,
category VARCHAR (64) NOT NULL,
town VARCHAR (64) NOT NULL,
date DATE ,
price INT ,
PRIMARY KEY (id)
);

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