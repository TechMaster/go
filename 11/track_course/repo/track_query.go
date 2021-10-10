package repo

import "track_course/model"

func Find_track_master_by_name(name string) (track_master *model.TrackMaster, err error) {
	track_master = new(model.TrackMaster)
	err = DB.Model(track_master).Where("name ILIKE ? AND ", "%"+name+"%").Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return track_master, nil
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


func 
