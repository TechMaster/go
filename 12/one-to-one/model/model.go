package model

type Person struct {
	Id       string `gorm:"primary_key"`
	FullName string
	Email    string
	Card     Card `gorm:"foreignKey:PersonId"`
}

type Card struct {
	Id       string `gorm:"primary_key"`
	Seri     string
	PersonId string
}
