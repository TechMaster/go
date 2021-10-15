package model

type Bar struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	FooId string `gorm:"column:foo_id"`
	Foo   Foo
}
