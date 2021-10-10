DROP TABLE IF EXISTS test.track_course;

CREATE TABLE test.track_course (
  track_id text REFERENCES test.track(id) ON DELETE CASCADE,
  course_id text REFERENCES test.course(id) ON DELETE CASCADE,
  display_order int NOT NULL -- Thứ tự của course trong track
)