DROP TYPE auth.rolenum;

CREATE TYPE auth.rolenum AS ENUM (
	'ADMIN',
	'STUDENT',
	'TRAINER',
	'SALE',
	'EMPLOYER',
	'AUTHOR',
	'EDITOR',
	'MAINTAINER');


DROP TABLE IF EXISTS auth.users;
CREATE TABLE users (
	id text PRIMARY KEY,
	name text NOT NULL,
	email text,
	mobile text,
	int_roles int[],
	enum_roles rolenum[]
);

DROP TABLE IF EXISTS auth.roles;
CREATE TABLE auth.roles (
	id int PRIMARY KEY,
	name text NOT NULL UNIQUE
);
INSERT INTO roles (id, name) VALUES (1, 'ADMIN');
INSERT INTO roles (id, name) VALUES (2, 'STUDENT');
INSERT INTO roles (id, name) VALUES (3, 'TRAINER');
INSERT INTO roles (id, name) VALUES (4, 'SALE');
INSERT INTO roles (id, name) VALUES (5, 'EMPLOYER');
INSERT INTO roles (id, name) VALUES (6, 'AUTHOR');
INSERT INTO roles (id, name) VALUES (7, 'EDITOR');
INSERT INTO roles (id, name) VALUES (8, 'MAINTAINER');

DROP TABLE IF EXISTS auth.user_role;
CREATE TABLE auth.user_role (
   user_id text REFERENCES users(id),
   role_id int REFERENCES roles(id)
);

--CREATE INDEX user_idx ON user_role USING hash (user_id, role_id);
CREATE UNIQUE INDEX user_idx ON user_role (user_id, role_id);
