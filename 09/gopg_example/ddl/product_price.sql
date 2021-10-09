DROP TABLE IF EXISTS test.laptop;

CREATE TABLE test.laptop (
  id text PRIMARY KEY,
  name text NOT NULL,
  cpu text NOT NULL
);


DROP TABLE IF EXISTS test.bike;

CREATE TABLE test.bike (
  id text PRIMARY KEY,
  name text NOT NULL,
  volume int
);


DROP TABLE IF EXISTS test.price_history;
CREATE TABLE test.price_history (
  id text PRIMARY KEY,
  item_id text NOT null,
  type text NOT null,
  price int,
  created_date timestamp
)

