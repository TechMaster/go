package main

import (
	"fmt"
	"gopgexample/model"
	"gopgexample/repo"
)

func main() {
	users := []model.User{} //Khai báo một slice chứa các đối tượng model.User
	err := repo.DB.Model(&users).Select()

	if err != nil {
		fmt.Println(err)
	} else {
		for _, user := range users {
			fmt.Println(user.Name)
		}
	}
}
