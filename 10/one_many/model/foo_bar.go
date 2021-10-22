package model

//------ Quan hệ Một - Nhiều
type Foo struct { //Tương đương với Book
	tableName struct{} `pg:"test.foo"` //Giống Annotation trong Java MemberShip --> member_ship
	Id        string   `pg:"id,pk"`
	Name      string
	Bars      []Bar `pg:"rel:has-many"`
}

type Bar struct { //Tương đương với Page
	tableName struct{} `pg:"test.bar"`
	Id        string   `pg:"id,pk"`
	Name      string
	FooId     string //Tự động chuyển thành foo_id
	Foo       Foo    `pg:"rel:has-one"`
}
