package person

import "fmt"

func BuildPerson() {
	address := &Address{
		City: "Hà Nội",
		Stress: "Nguyễn Trãi",
	}
	// Tạo lớp
	person := Person {
        Name: "Nguyễn Văn A",
		Email: "a@gmail.com",
		Age: 23,
		Address: address,
    }

	fmt.Println("Thông tin ban đầu")
    person.Print()

    personClone := person.Clone()

	fmt.Println("Thông tin sau khi clone")
    personClone.Print()

    // Thay đổi thông tin địa chỉ
	personClone.ChangeAddress("Lê Văn Lương")

	 //In lại thông tin person clone sau khi thay đổi thông tin
	fmt.Println("Thông tin sau thay đổi thông tin")
    personClone.Print()

    // In lại thông tin ban đầu
	fmt.Println("Thông tin ban đầu")
    person.Print()


	// Clone 2
	personClone02 := personClone.Clone()
	personClone02.Print()
	personClone02.ChangeAddress("Phạm Hùng")
	personClone02.Print()
	personClone.Print()
	person.Print()
}