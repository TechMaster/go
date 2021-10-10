package repo

/*
Admin, Sales, Maintainer có thể thay đổi track, course mà không cập nhật version.
Chỉ khi nào họ chủ định tạo ra một phiên bản mới new version mới
thì phiên bản track mới, course mới sẽ đươc tạo ra clone từ track mới nhất
*/

func Clone_track_to_new_version(track_master_id string) (err error) {
	return nil
}
