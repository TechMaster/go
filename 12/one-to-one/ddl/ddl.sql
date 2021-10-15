CREATE TABLE db.person (
  id varchar(10) PRIMARY KEY,
  full_name varchar(50),
  email varchar(50),
)

CREATE TABLE db.card (
  id varchar(10) PRIMARY KEY,
  seri varchar(50)
  person_id text REFERENCES db.person(id) ON DELETE CASCADE
)

INSERT INTO db.person(id,full_name,email) VALUES 
('1','anna','anna@gmail.com'),
('2','bob','bob@gmail.com'),
('3','john','john@gmail.com'),
('4','alice','alice@gmail.com')

INSERT INTO db.card(id,seri,person_id) VALUES 
('1','1234','1'),
('2','2345','2'),
('3','3456','3'),
('4','4567','4')