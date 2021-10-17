package model

import (
	"fmt"
	"time"
)

type User struct {
	tableName struct{}  `pg:"test.users"`
	Id        string    `pg:"id,pk" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Posts     []Post    `pg:"rel:has-many"`
}

/*
Struct này dùng để create và update user
*/
type CreateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

/*
In ra thông tin của user
*/
func (u *User) Print() {
	fmt.Println("Id : ", u.Id)
	fmt.Println("FullName : ", u.FullName)
	fmt.Println("Email : ", u.Email)
	fmt.Println("Phone : ", u.Phone)
	fmt.Println("CreatedAt : ", u.CreatedAt)
	fmt.Println("UpdatedAt : ", u.UpdatedAt)
	fmt.Println()
}
