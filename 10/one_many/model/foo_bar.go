package model

//------ Quan hệ Một - Nhiều
type Foo struct {
	tableName struct{} `pg:"test.foo"`
	Id        string   `pg:"id,pk"`
	Name      string
	Bars      []Bar `pg:"rel:has-many"`
}

type Bar struct {
	tableName struct{} `pg:"test.bar"`
	Id        string   `pg:"id,pk"`
	Name      string
	FooId     string //Tự động chuyển thành foo_id
	Foo       Foo    `pg:"rel:has-one"`
}
