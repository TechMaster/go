DROP TABLE IF EXISTS test.outline
CREATE TABLE test.outline {
  id text PRIMARY KEY,
	name text NOT NULL,
   status int -- draft: 0, active: 1, hide: 2. Tác giả outline hoặc admin, sale thay đối
}

DROP TABLE IF EXISTS test.section
CREATE TABLE test.section {
  id text PRIMARY KEY,
  outline_id text REFERENCES test.outline(id),
	name text NOT NULL,
  status int -- draft: 0, active: 1, hide: 2. Tác giả outline set
}

DROP TABLE IF EXISTS test.section_lecture
CREATE TABLE test.section {
  section_id text REFERENCES test.section(id),
  lecture_id text REFERENCES test.lecture(id),
  status int -- draft, active, hide. Tác giả outline set
}

DROP TABLE IF EXISTS test.lecture
CREATE TABLE test.lecture {
  id text PK,
  name text,
  status int -- draft, active, hide. Tác giả lecure có quyền thay đổi
}