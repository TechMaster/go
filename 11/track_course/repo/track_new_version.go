package repo

/*
Admin, Sales, Maintainer có thể thay đổi track, course mà không cập nhật version.
Chỉ khi nào thực sự cần thiết họ mới tạo phiên bản mới của track
Truyền vào track_master_id, tìm ra bản ghi có is_master = true, lấy được id của track, gọi là source_track_id
Tạo một track mới có id mới gọi là target_track_id

Select trong track_course có track_id bằng source_track_id
Tạo mới tất cả các bản ghi này, thay source_track_id bằng target_track_id
*/

func Clone_track_to_new_version(track_master_id string) (err error) {
	return nil
}
