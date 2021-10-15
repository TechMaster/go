package repo

import "basic/model"

func DemoDelete() {
	// Khởi tạo transaction
	tx := DB.Begin()
	if err := tx.Error; err != nil {
		return
	}

	// Xóa bản ghi
	std := model.Student{ Id: "1" }
	tx.Delete(&std)

	// Xóa bản ghi theo điều kiện
	tx.Where("full_name = ?", "bob").Delete(&model.Student{})

	// Xóa theo khóa chính
	tx.Delete(&model.Student{}, "2")

	// Xóa theo danh sách khóa chính
	tx.Delete(&model.Student{}, []string{"3", "4"})

	tx.Rollback()
}