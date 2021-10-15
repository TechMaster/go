package repo

import "basic/model"

func DemoSelect() {
	var std model.Student

	//  Lấy bản ghi đầu tiên sau khi sắp xếp theo khóa chính
	// DB.First(&std)

	//  Lấy bản ghi đầu tiên
	// DB.Take(&std)

	//  Lấy bản ghi cuối cùng sau khi sắp xếp theo khóa chính
	DB.Last(&std)

	// Tìm kiếm theo khóa chính
	DB.First(&model.Student{}, "2")
	DB.First(&model.Student{}, "id = ?", "2")

	// Tìm kiếm theo danh sách khóa chính
	DB.Find(&[]model.Student{}, []string{"2", "3", "4"})

	// Lấy tất cả bản ghi trong bảng
	var students []model.Student
	DB.Find(&students)

	// Condition
	DB.Where("full_name = ?", "bob").First(&model.Student{})
	
}
