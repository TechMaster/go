DROP TABLE IF EXISTS db.member;
CREATE TABLE db.member (
  id VARCHAR(10) PRIMARY KEY,
  name VARCHAR(30) NOT NULL 
);

DROP TABLE IF EXISTS db.club;
CREATE TABLE db.club (
  id VARCHAR(10) PRIMARY KEY,
  name VARCHAR(30) NOT NULL 
);

DROP TABLE IF EXISTS db.member_club;
CREATE TABLE db.member_club (
  member_id VARCHAR(10) REFERENCES db.member(id) ON DELETE cascade,
  club_id VARCHAR(10) REFERENCES db.club(id) ON DELETE cascade,
  active BOOLEAN NOT NULL
);