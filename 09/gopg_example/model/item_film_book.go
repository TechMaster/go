package model

type Item struct {
	tableName struct{} `pg:"_"` // "_" means no name
	Id        string   `pg:"id,pk"`
	Name      string
	Price     int
}

type Book struct {
	tableName struct{} `pg:"test.book"`
	Item      `pg:",inherit"`
	Pages     int
}

type Film struct {
	tableName struct{} `pg:"test.film"`
	Item      `pg:",inherit"`
	Director  string
	Actors    []Actor `pg:"rel:has-many"`
}

type Actor struct {
	tableName struct{} `pg:"test.actor"`
	Id        string   `pg:"id,pk"`
	Name      string
	FilmId    string
	Film      Film `pg:"rel:has-one"`
}
