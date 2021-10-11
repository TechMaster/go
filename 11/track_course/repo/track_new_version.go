package repo

import (
	"time"
	"track_course/model"
)

/*
Admin, Sales, Maintainer có thể thay đổi track, course mà không cập nhật version.
Chỉ khi nào thực sự cần thiết họ mới tạo phiên bản mới của track
Truyền vào track_master_id, tìm ra bản ghi có is_master = true, lấy được id của track, gọi là source_track_id
Tạo một track mới có id mới gọi là target_track_id

Select trong track_course có track_id bằng source_track_id
Tạo mới tất cả các bản ghi này, thay source_track_id bằng target_track_id
*/

func Clone_track_to_new_version(track_master_id string) (track *model.Track, err error) {
	//Lấy source_track_id
	var source_track_id string

	err = DB.Model(&source_track_id).Column("id").
		Where("track.master_id = ? AND track.is_master = 'true'", track_master_id).
		Select()

	//var track_name = Get_track_name(track_master_id)

	//Chuyển is_master của tất cả track về false
	track = &model.Track{
		IsMaster: false,
	}
	_, err = DB.Model(track).Column("is_master").
		Where("track.master_id = ? ", track_master_id).
		Update()

	//Tạo track mới có track_master_id giữ nguyên
	track = &model.Track{
		Id:          NewID(),
		MasterId:    track_master_id,
		IsMaster:    true,
		Status:      model.STATUS_ACTIVE,
		Version:     0,
		Name:        "test",
		Description: "java",
		BasePrice:   0,
		Price:       0,
		Lessons:     0,
		CreatedDate: time.Now().Add(time.Hour),
	}

	_, err = DB.Model(track).Insert()

	if err != nil {
		return nil, err
	}
	return track, nil

}

// func Get_track_name(track_master_id string) string {

// 	var track_name string

// 	_, err := DB.Query(&track_name,
// 		"SELECT track.name FROM test.track AS track WHERE track.master_id = ? AND is_master = 'true'", track_master_id)
// 	// _, err := DB.Model(track_name).Column("name").
// 	// 	Where("master_id = ? AND is_master = 'true'", track_master_id).
// 	// 	Select()
// 	if err != nil {
// 		return ""
// 	}
// 	return track_name
// }
