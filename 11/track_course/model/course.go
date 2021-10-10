package model

import "time"

type Course struct {
	tableName   struct{} `pg:"course.track"`
	Id          string   `pg:"id,pk"`
	MasterId    string
	IsMaster    bool `pg:",use_zero"` //Bắt buộc phải dùng pg:",use_zero" nếu không giá trị false biến thành null
	Version     int  `pg:",use_zero"`
	Name        string
	Description string
	BasePrice   int       `pg:",use_zero"`
	Price       int       `pg:",use_zero"`
	Lessons     int       `pg:",use_zero"`
	CreatedDate time.Time `pg:"type:timestamp without time zone,default:now()"`
}
