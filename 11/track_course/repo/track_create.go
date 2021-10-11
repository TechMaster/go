package repo

import (
	"time"
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
		CreatedDate: time.Now().Add(time.Hour),
	}
	_, err = DB.Model(track).Insert()
	if err != nil {
		return nil, err
	} else {
		return track, nil
	}
}


<<<<<<< HEAD
/*
Clone từ một track có sẵn ra một trac
*/
func Clone_new_version_of_track(track_id string) (err error) {
	return nil
}

func generateRandomTrackId() string {
	var ids []string
	_, err := model.DB.Query(&ids, `
		SELECT id FROM test.tracl
	`)
	if err != nil {
		return ""
	}
	return ids[random.Intn(len(ids))]
}
=======

>>>>>>> 89e4cb12b724c5aed9bce5cc1cd0c2eff07416f1
