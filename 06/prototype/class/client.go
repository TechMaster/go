package class

import "fmt"

func BuildClass() {
	// Tạo lớp
	class := Class {
        Name: "Golang One Mount",
		Session: 21,
		Teacher: "Trịnh Minh Cường",
    }

	// In ra thông tin lớp
	fmt.Println("Thông tin lớp ban đầu")
    class.Print() 

    // Clone 1 lớp khác từ lớp đã tạo trước đó
    classClone := class.Clone()

	// In thông tin lớp vừa clone
	fmt.Println("Thông tin lớp sau khi clone")
    classClone.Print()

    // Thay đổi thông tin của lớp vừa clone
    classClone.SetName("Java One Mount") // Đổi tên lớp
    classClone.SetSession(30) // Đổi số buổi học
	classClone.SetTeacher("Lục Thanh Ngọc") // Đổi tên giảng viên

	 //In lại thông tin lớp clone sau khi thay đổi thông tin
	fmt.Println("Thông tin lớp sau khi thay đổi thông tin")
    classClone.Print()

    // In lại thông tin lớp ban đầu
	fmt.Println("Thông tin lớp ban đầu")
    class.Print()
}