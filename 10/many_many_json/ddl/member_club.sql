DROP TABLE IF EXISTS test.memberj;
CREATE TABLE test.memberj (
  id text PRIMARY key,
  name text NOT NULL,
  memberships jsonb
);

DROP TABLE IF EXISTS test.clubj;
CREATE TABLE test.clubj (
  id text PRIMARY key,
  name text NOT NULL,
  memberships jsonb
);