package repo

import "track_course/model"

func Find_track_by_name(name string) (track *model.Track, err error) {

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
