package model

type Student struct {
	Id       string `gorm:"primary_key"`
	FullName string
	Email    string
	Phone    string
	CardId   string
}
