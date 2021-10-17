DROP TABLE IF EXISTS db.users
CREATE TABLE db.users (
	id varchar(10) PRIMARY KEY,
	full_name varchar(50),
	email varchar(50),
	phone varchar(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)

DROP TABLE IF EXISTS tdbest.posts
CREATE TABLE db.posts (
	id varchar(10) PRIMARY KEY,
	title varchar(500),
	content varchar(1000),
	author_id varchar(10) REFERENCES db.users(id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)