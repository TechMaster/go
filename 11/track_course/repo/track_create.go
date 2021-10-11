package repo

import (
	"track_course/model"
)

/*
Tạo track mới hoàn toàn
*/
func Create_new_track(name string, desc string) (track *model.Track, err error) {
	track = &model.Track{
		Id:          NewID(),
		MasterId:    NewID(),
		IsMaster:    true,
		Status:      model.STATUS_ACTIVE,
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



