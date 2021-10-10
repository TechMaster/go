/*
Lưu thông tin khoá học mới nhất
*/
DROP TABLE IF EXISTS test.course_master;
CREATE TABLE test.course_master (
	id text PRIMARY KEY, -- 5 ký tự sinh ngẫu nhiên nhờ go-nanoid
	name text NOT NULL UNIQUE,  -- Tên cập nhật nhất
	description text NOT NULL,
  lessons int NOT NULL, -- số lượng lesson trong khoá học
	base_price int NOT NULL, -- giá tính toán
	price int NOT NULL -- giá đề xuất sẽ hiển thị lên web site  
);


/*
Lưu lịch sử các phiên bản thay đổi lớn của một khoá học
*/
DROP TABLE IF EXISTS test.course;
CREATE TABLE test.course (
	id text PRIMARY KEY, -- Kết hợp master_id + version. Dài 8 ký tự
	master_id text REFERENCES test.course_master(id),
	version text NOT NULL, --Chạy từ `000` đến `999`
	name text NOT NULL UNIQUE
	description text NOT NULL,
  lessons int NOT NULL, -- số lượng lesson trong khoá học
	base_price int NOT NULL, -- giá tính toán
	price int NOT NULL, -- giá đề xuất sẽ hiển thị lên web site
	created_date timestamp without time zone default (now()) -- ngày tạo phiên bản này
);