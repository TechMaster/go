package model

type Bar struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	FooId string `gorm:"column:foo_id"`
	Foo   Foo
}

type Foo struct {
	Id   string `gorm:"primaryKey"`
	Name string
	Bars []Bar `gorm:"foreignKey:FooId"`
}

func(b *Bar) TableName() string {
	return "bar"
}

func(b *Foo) TableName() string {
	return "foo"
}
