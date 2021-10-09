DROP TABLE IF EXISTS test.foo;
CREATE TABLE test.foo (
  id text PRIMARY KEY,
  name text NOT NULL 
);


DROP TABLE IF EXISTS test.bar;
CREATE TABLE test.bar (
  id text PRIMARY KEY,
  name text NOT null,
  foo_id text REFERENCES test.foo(id) ON DELETE CASCADE
);

INSERT INTO test.foo (id, name) VALUES ('ox-01', 'foo1');
INSERT INTO test.foo (id, name) VALUES ('ox-02', 'foo2');

INSERT INTO test.bar (id, name, foo_id) VALUES ('bar1', 'this is bar1', 'ox-01');
INSERT INTO test.bar (id, name, foo_id) VALUES ('bar2', 'this is bar 2', 'ox-02');
INSERT INTO test.bar (id, name, foo_id) VALUES ('bar3', 'this is bar 3', 'ox-01');
INSERT INTO test.bar (id, name, foo_id) VALUES ('bar4', 'this is bar 4', 'ox-01');