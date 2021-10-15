package repo

import (
	"basic/model"
)

func DemoUpdate() {
	// Khởi tạo transaction
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		return
	}

	// Sử dụng Save
	var std model.Student
	tx.First(&std)

	std.FullName = "Nguyễn Văn D"
	std.Email = "a@gmail.com"
	tx.Save(&std)

	// Update single column
	tx.Model(&std).Update("full_name", "hello-1")

	// Sử dụng condition
	tx.Model(&model.Student{}).Where("email = ?", "bob@gmail.com").Update("full_name", "bob1")

	// Updates multiple columns
	tx.Model(&model.Student{}).Where("email = ?", "bob@gmail.com").Updates(model.Student{FullName: "bob1", Email: "098888888"})

	tx.Rollback()
}
