package model

import "time"

type TrackMaster struct {
	tableName   struct{} `pg:"test.track_master"`
	Id          string   `pg:"id,pk"`
	Version     int      //Chỉ lấy version mới có track active, bỏ qua version track draft
	Name        string
	Description string
	BasePrice   int
	Price       int
	Lessons     int
	Status      string
	Tracks      []Track `pg:"rel:has-many"`
}

type Track struct {
	tableName     struct{}    `pg:"test.track"`
	Id            string      `pg:"id,pk"`
	TrackMasterId string      `pg:"master_id"`
	TrackMaster   TrackMaster `pg:"rel:has-one"`
	Version       int
	Name          string
	Description   string
	BasePrice     int
	Price         int
	Lessons       int
	CreatedDate   time.Time `pg:"type:timestamp without time zone,default:now()"`
}
