package repo

/*
Chỉ thêm  course vào track có version mới nhất, thường là k
và cũng chỉ nên dùng course có version mới nhất
*/
func Add_course_to_track(track_id string, course_id string) (err error) {
	return nil
}

/*
Chỉ xoá course khỏi track có version mới nhất.
Nếu xoá course khỏi track đã bán cho sinh viên, họ sẽ phàn nàn
và cũng chỉ nên dùng course có version mới nhất
*/
func Remove_course_from_track(track_id string, course_id string) (err error) {
	return nil
}

/*
thay đổi vị trí của course trong track
*/
func Change_order_course_in_track(track_id string, course_id string, display_order int) (err error) {
	return nil
}
