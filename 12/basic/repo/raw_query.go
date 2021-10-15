package repo

import (
	"basic/model"
	"database/sql"
)

func DemoRawQuery() {
	// Khởi tạo transaction
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		return
	}

	// Lấy student theo id
	var std model.Student
	tx.Raw("SELECT id, full_name, email FROM student WHERE id = ?", "1").Scan(&std)

	// Lấy danh sách student
	var students []model.Student
	tx.Raw("SELECT id, full_name, email FROM student").Scan(&students)

	// Lấy 2 người đầu tiên trong danh sách
	tx.Raw("SELECT id, full_name, email FROM student LIMIT 2").Scan(&students)

	// Update name của user có id = 1
	tx.Exec("UPDATE student set full_name = ? WHERE id = ?", "abc", "1")

	// Name Argument
	tx.Where("full_name = @name", sql.Named("name", "bob")).Find(&model.Student{})

	// Custom struct Argument
	type Custom struct {
		Name  string
		Email string
	}

	tx.Raw("SELECT * FROM student WHERE full_name = @Name AND email = @Email", Custom{ Name: "bob", Email: "bob@gmail.com"}).Find(&model.Student{})

	tx.Rollback()
}
