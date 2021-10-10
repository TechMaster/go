package model

import "time"

type Track struct {
	tableName   struct{} `pg:"test.track"`
	Id          string   `pg:"id,pk"`
	MasterId    string
	IsMaster    bool `pg:",use_zero"` //Bắt buộc phải dùng pg:",use_zero" nếu không giá trị false biến thành null
	Version     int
	Name        string
	Description string
	BasePrice   int
	Price       int
	Lessons     int
	CreatedDate time.Time `pg:"type:timestamp without time zone,default:now()"`
}
