package model

type Foo struct {
	Id   string `gorm:"primaryKey"`
	Name string
	Bars []Bar `gorm:"foreignKey:FooId"`
}
