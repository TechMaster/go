CREATE TABLE db.foo (
  id varchar(10) PRIMARY KEY,
  name varchar(50) NOT NULL 
);

CREATE TABLE db.bar (
  id varchar(10) PRIMARY KEY,
  name varchar(50) NOT null,
  foo_id varchar(10) REFERENCES db.foo(id) ON DELETE CASCADE
);

INSERT INTO db.foo (id, name) VALUES 
('ox-01', 'foo1'),
('ox-02', 'foo2')

INSERT INTO db.bar (id, name, foo_id) VALUES 
('bar1', 'this is bar 1', 'ox-01'),
('bar2', 'this is bar 2', 'ox-02'),
('bar3', 'this is bar 3', 'ox-01'),
('bar4', 'this is bar 4', 'ox-01')