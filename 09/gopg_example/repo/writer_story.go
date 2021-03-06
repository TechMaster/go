package repo

import (
	"fmt"
	"gopgdemo/model"
)

/*
Tạo bảng tạm Writer và Story quan hệ 1:1
*/

func Create_writer_story() (err error) {
	writer1 := &model.Writer{
		Name:   "admin",
		Emails: []string{"cuong@techmaster.vn", "minhcuong@gmail.com"},
	}
	_, err = DB.Model(writer1).Insert()
	if err != nil {
		return err
	}

	_, err = DB.Model(&model.Writer{
		Name:   "root",
		Emails: []string{"root@techmaster.vn", "root@gmail.com"},
	}).Insert()
	if err != nil {
		return err
	}

	story1 := &model.Story{
		Title:    "Cool story",
		AuthorId: writer1.Id,
	}
	_, err = DB.Model(story1).Insert()
	if err != nil {
		return err
	}

	// Select model.Writer by primary key.
	writer := &model.Writer{Id: writer1.Id}
	err = DB.Model(writer).WherePK().Select()
	if err != nil {
		return err
	}

	// Select all model.Writers.
	var writers []model.Writer
	err = DB.Model(&writers).Select()
	if err != nil {
		return err
	}

	// Select story and associated author in one query.
	story := new(model.Story)
	err = DB.Model(story).
		Relation("Author").
		Where("story.id = ?", story1.Id).
		Select()
	if err != nil {
		return err
	}

	fmt.Println(writer)
	fmt.Println(writers)
	fmt.Println(story)
	return nil
}
