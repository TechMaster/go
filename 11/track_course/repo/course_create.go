package repo

/*
Hãy code giống hệt như track_create
*/

// func Create_new_course(name string, desc string) (course *model.Course, err error) {
// 	course = &model.Course{
// 		Id:          NewID(),
// 		MasterId:    NewID(),
// 		IsMaster:    true,
// 		Status:      model.STATUS_ACTIVE,
// 		Lessons:     0,
// 		Version:     0,
// 		Name:        name,
// 		Description: desc,
// 		BasePrice:   0,
// 		Price:       0,
// 		CreatedDate: time.Now().Add(time.Hour),
// 	}

// 	_, err = DB.Model(course).Insert()

// 	if err != nil {
// 		return nil, err
// 	} else {
// 		return course, nil
// 	}
// }

func generateRandomCourseId() string {
	var ids []string
	_, err := DB.Query(&ids, `
		SELECT id FROM test.course
	`)
	if err != nil {
		return ""
	}
	return ids[random.Intn(len(ids))]
}
