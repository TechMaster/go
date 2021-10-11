package repo

import "track_course/model"

func Find_track_by_name(name string) (track *model.Track, err error) {
	track = new(model.Track)
	err = DB.Model(track).Where("name ILIKE ? AND ", "%"+name+"%").Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return track, nil
}

func Find_track_by_id(id string) (track *model.Track, err error) {
	track = &model.Track{
		Id: id,
	}
	err = DB.Model(track).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return track, nil
}

/*
Truyền vào bất kỳ track_id nào, lấy master_id rồi lấy max_version
Chú ý chỉ lấy những track có trạng thái active
lastes_track_id là track_id tương ứng version mới nhất
*/
func Get_max_version_track(track_id string) (max_version int, lastes_track_id string, err error) {
	return 0, "", nil
}
