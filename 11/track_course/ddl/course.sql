
/*
Lưu lịch sử các phiên bản thay đổi lớn của một khoá học
*/
DROP TABLE IF EXISTS test.course;
CREATE TABLE test.course (
	id text PRIMARY KEY, -- Kết hợp master_id + version. Dài 8 ký tự
	master_id text,
	is_master  bool NOT NULL,
	version int,
	name text NOT NULL UNIQUE
	description text NOT NULL,
  lessons int NOT NULL, -- số lượng lesson trong khoá học
	base_price int NOT NULL, -- giá tính toán
	price int NOT NULL, -- giá đề xuất sẽ hiển thị lên web site
	created_date timestamp without time zone default (now()) -- ngày tạo phiên bản này
);