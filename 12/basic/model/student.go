package model

import "fmt"

/*
GORM sẽ tự động chuyển tên struct thành tên bảng theo dạng là snake_case
Ví dụ : Student -> students, User -> users, UserDetail -> user_details
*/
type Student struct {
	Id       string `gorm:"primary_key"` // Id là trường mặc định được sử dụng làm khóa chính
	FullName string
	Email    string
	Phone    string
	CardId   string
}

func(s *Student) ToString() {
	fmt.Println("Id : ", s.Id)
	fmt.Println("FullName : ", s.FullName)
	fmt.Println("Email : ", s.Email)
	fmt.Println("Phone : ", s.Phone)
	fmt.Println("CardId : ", s.CardId)
}

/*
Chúng ta có thể thay đổi tên bảng default bằng cách implement interface Tabler

type Tabler interface {
	TableName() string
}
*/

func(s *Student) TableName() string{
	return "student"
}
