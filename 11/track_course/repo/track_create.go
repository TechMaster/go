package repo

import (
	"track_course/model"
)

/*
Tạo track_master và phiên bản track đầu tiên version "000"
*/
func Create_new_track_master(name string, desc string) (track_master *model.TrackMaster, err error) {
	track_master = &model.TrackMaster{
		Id:          NewID(5),
		Version:     "000",
		Name:        name,
		Description: desc,
		BasePrice:   0,
		Price:       0,
		Lessons:     0,
	}

	track := &model.Track{
		Id:            track_master.Id + track_master.Version,
		TrackMasterId: track_master.Id,
		Version:       track_master.Version,
		Name:          name,
		Description:   desc,
		BasePrice:     0,
		Price:         0,
		Lessons:       0,
	}

	transaction, err := DB.Begin()
	if err != nil {
		return nil, err
	}

	_, err = transaction.Model(track_master).Insert()
	if !check_err(err, transaction) {
		return nil, err
	}

	_, err = transaction.Model(track).Insert()
	if !check_err(err, transaction) {
		return nil, err
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	} else {
		return track_master, nil
	}
}
