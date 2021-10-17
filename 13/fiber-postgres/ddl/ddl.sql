DROP TABLE IF EXISTS test.users
CREATE TABLE test.users (
	id text PRIMARY KEY,
	full_name text,
	email text,
	phone text,
    created_at TIMESTAMP with time zone,
    updated_at TIMESTAMP with time zone
)

DROP TABLE IF EXISTS test.posts
CREATE TABLE test.posts (
	id text PRIMARY KEY,
	title text,
	content text,
	author_id text REFERENCES test.users(id) ON DELETE CASCADE,
    created_at TIMESTAMP with time zone,
    updated_at TIMESTAMP with time zone
)