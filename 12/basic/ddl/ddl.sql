-- db.student definition
CREATE TABLE db.student (
  id varchar(10) NOT NULL,
  full_name varchar(50),
  email varchar(50),
  phone varchar(50),
  card_id varchar(50),
  PRIMARY KEY (id)
)

INSERT INTO db.student(id,full_name,email,phone,card_id) VALUES 
('1','anna','anna@gmail.com','0987654321','100'),
('2','bob','bob@gmail.com','0344005934','101'),
('3','john','john@gmail.com','0964543470','102'),
('4','alice','alice@gmail.com','0987123456','103')