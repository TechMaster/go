DROP TABLE IF EXISTS test.price;

CREATE TABLE test.price (
  id serial PRIMARY KEY,
  item_id text NOT NULL,
  discount_amount int default 0,
  created_date timestamp NOT NULL
)