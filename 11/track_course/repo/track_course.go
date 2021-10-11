package repo

import "track_course/model"

func Create_track_course() (track_course *model.Track_Course, err error) {

	track_course = &model.Track_Course{
		TrackId:      generateRandomCourseId(),
		CourseId:     generateRandomCourseId(),
		DisplayOrder: 1,
	}

	_, err = DB.Model(track_course).Insert()
	if err != nil {
		return nil, err
	} else {
		return track_course, nil
	}

}
