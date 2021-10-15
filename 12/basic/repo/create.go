package repo

import (
	"basic/model"
	"basic/util"

	"github.com/brianvoe/gofakeit/v6"
)

func DemoCreate() {
	// Khởi tạo transaction
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		return
	}

	// Tạo đối tượng
	std := model.Student{
		Id:       util.NewID(),
		FullName: gofakeit.Animal(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		CardId:   util.NewID(),
	}
	tx.Create(&std)

	// Tạo đối tượng với các trường được chỉ định
	tx.Select("Id", "FullName", "Email", "CardId").Create(&model.Student{
		FullName: gofakeit.Animal(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		CardId:   util.NewID(),
	})

	tx.Omit("Email", "Phone", "CardId").Create(&model.Student{
		Id:       util.NewID(),
		FullName: gofakeit.Animal(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		CardId:   util.NewID(),
	})

	// Tạo nhiều đối tượng cùng lúc
	var students []model.Student
	for i := 0; i < 5; i++ {
		newStudent := model.Student{
			Id:       util.NewID(),
			FullName: gofakeit.Animal(),
			Email:    gofakeit.Email(),
			Phone:    gofakeit.Phone(),
			CardId:   util.NewID(),
		}
		students = append(students, newStudent)
	}

	tx.Create(&students)

	// Tạo đối tượng với struct
	DB.Model(&model.Student{}).Create(map[string]interface{}{
		"Id":       util.NewID(),
		"FullName": gofakeit.Animal(),
		"Email":    gofakeit.Email(),
	})

	tx.Rollback()
}
