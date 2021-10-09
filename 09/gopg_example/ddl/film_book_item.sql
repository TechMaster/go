DROP TABLE IF EXISTS test.item;

CREATE TABLE test.item (
  id text PRIMARY KEY,
  name text NOT NULL,
  price int NOT NULL
);

DROP TABLE IF EXISTS test.book;

CREATE TABLE test.book (
  pages int
) INHERITS (test.item);

DROP TABLE IF EXISTS test.film;

CREATE TABLE test.film (
  director text
) INHERITS (test.item);

DROP TABLE IF EXISTS test.actor;

CREATE TABLE test.actor (
  id text PRIMARY KEY,
  name text NOT NULL,
  film_id text
);
