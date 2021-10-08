package model

import "fmt"

type Writer struct {
	Id     int64
	Name   string
	Emails []string
}

func (u Writer) String() string {
	return fmt.Sprintf("Writer<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *Writer `pg:"rel:has-one"` //Quan hệ 1:1 một tác giả viết một truyện
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}
