/*
Lưu thông tin lộ trình mới nhất
*/
DROP TABLE IF EXISTS test.track_master;
CREATE TABLE test.track_master (
	id text PRIMARY KEY, -- 5 ký tự sinh ngẫu nhiên nhờ go-nanoid
	name text NOT NULL UNIQUE,  -- Tên cập nhật nhất
	description text NOT NULL,
	base_price int, -- giá tính toán
	price int , -- giá đề xuất sẽ hiển thị lên web site
	lessons int, -- số lượng lesson = tổng các khoá học thành phần
	version text --phiên bản cuối cùng
);


/*
Lưu lịch sử các phiên bản thay đổi lớn của một lộ trình
*/
DROP TABLE IF EXISTS test.track;
CREATE TABLE test.track (
	id text PRIMARY KEY, -- Kết hợp master_id + version. Dài 8 ký tự
	master_id text REFERENCES test.track_master(id),
	version text NOT NULL, --Chạy từ `000` đến `999`
	name text NOT NULL UNIQUE,
	description text NOT NULL,
	base_price int, -- gi	á tính toán
	price int, -- giá đề xuất sẽ hiển thị lên web site
	lessons int, -- số lượng lesson = tổng các khoá học thành phần
	created_date timestamp without time zone default (now()) -- ngày tạo phiên bản này
);