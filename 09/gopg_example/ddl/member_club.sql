DROP TABLE IF EXISTS test.member;
CREATE TABLE test.member (
  id text PRIMARY key,
  name text NOT NULL 
);

DROP TABLE IF EXISTS test.club;
CREATE TABLE test.club (
  id text PRIMARY key,
  name text NOT NULL 
);

DROP TABLE IF EXISTS test.member_club;
CREATE TABLE test.member_club (
  member_id text REFERENCES test.member(id) ON DELETE cascade,
  club_id text REFERENCES test.club(id) ON DELETE cascade,
  active boolean NOT NULL
);