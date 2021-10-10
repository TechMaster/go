package repo

import (
	"track_course/model"
)

/*
Tạo track_master và phiên bản track đầu tiên version "000"
*/
func Create_new_track(name string, desc string) (track *model.Track, err error) {
	track = &model.Track{
		Id:          NewID(),
		MasterId:    NewID(),
		IsMaster:    true,
		Version:     0,
		Name:        name,
		Description: desc,
		BasePrice:   0,
		Price:       0,
		Lessons:     0,
	}
	_, err = DB.Model(track).Insert()
	if err != nil {
		return nil, err
	} else {
		return track, nil
	}
}

/*
Truyền vào bất kỳ track_id nào, lấy master_id rồi lấy max_version
*/
func Get_max_version_track(track_id string) (max_version int, err error) {
	return 0, nil
}
func Clone_new_version_of_track(track_id string) (err error) {
	return nil
}
