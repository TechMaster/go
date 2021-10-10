DROP TYPE IF EXISTS test.rolenum;

CREATE TYPE test.rolenum AS ENUM (
'ADMIN',
'STUDENT',
'TRAINER',
'SALE',
'EMPLOYER',
'AUTHOR',
'EDITOR',
'MAINTAINER');


DROP TABLE IF EXISTS test.users;
CREATE TABLE test.users (
	id text PRIMARY KEY,
	name text NOT NULL,
	email text,
	mobile text,
	int_roles int[],
	enum_roles test.rolenum[]
);

DROP TABLE IF EXISTS test.roles;
CREATE TABLE test.roles (
	id int PRIMARY KEY,
	name text NOT NULL UNIQUE
);
INSERT INTO test.roles (id, name) VALUES (1, 'ADMIN');
INSERT INTO test.roles (id, name) VALUES (2, 'STUDENT');
INSERT INTO test.roles (id, name) VALUES (3, 'TRAINER');
INSERT INTO test.roles (id, name) VALUES (4, 'SALE');
INSERT INTO test.roles (id, name) VALUES (5, 'EMPLOYER');
INSERT INTO test.roles (id, name) VALUES (6, 'AUTHOR');
INSERT INTO test.roles (id, name) VALUES (7, 'EDITOR');
INSERT INTO test.roles (id, name) VALUES (8, 'MAINTAINER');

DROP TABLE IF EXISTS test.user_role;
CREATE TABLE test.user_role (
   user_id text REFERENCES test.users(id),
   role_id int REFERENCES test.roles(id)
);

CREATE UNIQUE INDEX user_idx ON test.user_role (user_id, role_id);