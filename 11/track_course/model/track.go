package model

import "time"

type Track struct {
	tableName   struct{} `pg:"test.track"`
	Id          string   `pg:"id,pk"`
	MasterId    string
	IsMaster    bool `pg:",use_zero"` //Bắt buộc phải dùng pg:",use_zero" nếu không giá trị false biến thành null
	Version     int  `pg:",use_zero"`
	Name        string
	Description string
	Status      int       `pg:",use_zero"`
	BasePrice   int       `pg:",use_zero"`
	Price       int       `pg:",use_zero"`
	Lessons     int       `pg:",use_zero"`
	CreatedDate time.Time `pg:"type:timestamp without time zone,default:now()"`
}
