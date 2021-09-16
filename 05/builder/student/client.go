package student

import "fmt"

func BuildStudent() {
	student := newStudentBuilder().setName("Nguyễn Văn A").setSex("Nam").setAge(20).getStudent()
	printDetail(student)
}

func printDetail(s student) {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Sex: ", s.Sex)
	fmt.Println("Age: ", s.Age)
}