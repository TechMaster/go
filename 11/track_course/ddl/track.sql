
/*
Lưu lịch sử các phiên bản thay đổi lớn của một lộ trình
*/
DROP TABLE IF EXISTS test.track;
CREATE TABLE test.track (
	id text PRIMARY KEY,
	master_id text,
	is_master  bool NOT NULL,
	version int,
	name text NOT NULL,
	description text NOT NULL,
	status int NOT NULL,
	base_price int, -- giá tính toán theo số buổi x học phí
	price int, -- giá đề xuất sẽ hiển thị lên web site
	lessons int, -- số lượng lesson = tổng các khoá học thành phần
	created_date timestamp without time zone default (now()) -- ngày tạo phiên bản này
);