package model

type Track struct {
	tableName struct{} `pg:"content.track"`
	Id        string   `pg:",pk"`
	Name      string
	Courses   []Course `pg:"many2many:content.track_course"`
}

type Course struct {
	tableName struct{} `pg:"content.course"`
	Id        string   `pg:",pk"`
	Name      string
	Tracks    []Track `pg:"many2many:content.track_course"`
}

type Track_Course struct {
	tableName    struct{} `pg:"content.track_course"`
	Track_id     string
	Course_id    string
	DisplayOrder int
}
