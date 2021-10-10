package model

import "time"

type Course struct {
	tableName   struct{} `pg:"test.course"`
	Id          string   `pg:"id,pk"`
	MasterId    string
	IsMaster    bool `pg:",use_zero"`
	Version     int  `pg:",use_zero"`
	Name        string
	Description string
	Lessons     int       `pg:",use_zero"`
	BasePrice   int       `pg:",use_zero"`
	Price       int       `pg:",use_zero"`
	CreatedDate time.Time `pg:"type:timestamp without time zone,default:now()"`
}
