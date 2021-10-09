DROP TYPE IF EXISTS content.status;
CREATE TYPE status AS ENUM ('draft', 'active', 'hidden');

DROP TABLE IF EXISTS content.track;
CREATE TABLE content.track(
   id text PRIMARY KEY,
   name text NOT NULL,
   status status NOT NULL
);
CREATE INDEX idx_name ON content.track using hash(name);


DROP TABLE IF EXISTS content.course;
CREATE TABLE content.course(
   id text PRIMARY KEY,
   name text NOT NULL,
   status status,
   number_lesson int
);

DROP TABLE IF EXISTS content.track_course;
CREATE TABLE content.track_course(
   track_id text REFERENCES content.track(id),
   course_id text REFERENCES content.course(id)
);

/*
truncate table content.track;
truncate table content.outline cascade;
truncate table content.course;
*/

