package model

type Track_Course struct {
	tableName    struct{} `pg:"test.track_course"`
	TrackId      string   `pg:",pk"`
	CourseId     string   `pg:",pk"`
	DisplayOrder int      `pg:"default:0"`
}
