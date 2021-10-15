package repo

import (
	"basic/model"
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

	tx.Rollback()
}
