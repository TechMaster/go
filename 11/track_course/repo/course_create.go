package repo

/*
Hãy code giống hệt như track_create
*/

import (
	"time"
	"track_course/model"
)


func Create_new_course(name string, desc string) (course *model.Course, err error) {
	course = &model.Course{
		Id:          NewID(),
		MasterId:    NewID(),
		Lessons:     0,
		Version:     0,
		Name:        name,
		Description: desc,
		BasePrice:   0,
		Price:       0,
		CreatedDate: time.Now().Add(time.Hour),
	}

	_, err = DB.Model(course).Insert()

	if err != nil {
		return nil, err
	} else {
		return course, nil
	}
}
