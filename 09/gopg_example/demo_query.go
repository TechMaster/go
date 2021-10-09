package main

import (
	"fmt"
	"gopgdemo/model"
	"gopgdemo/repo"

	"github.com/brianvoe/gofakeit/v6"
)

func queryUsingModel() {
	users := []model.User{} //Khai báo một slice chứa các đối tượng model.User
	//Đây là lệnh truy vấn sử dụng Model
	err := repo.DB.Model(&users).Limit(10).Select()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, user := range users {
			fmt.Println(user.Name)
		}
	}
}

func queryPlainSQL() {
	users := []model.User{} //Khai báo một slice chứa các đối tượng model.User
	result, err := repo.DB.Query(&users, "SELECT u.id, u.name FROM auth.users u LIMIT 10")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Số dòng trả về ", result.RowsReturned())
		for _, user := range users {
			fmt.Println(user.Name)
		}
	}
}

func createUser() {
	repo.DB.Model(model.User{
		Id:   repo.NewID(),
		Name: gofakeit.Name(),
	})
}
