package main

import "fmt"

type Person struct {
	FirstName string //kiểu đứng sau tên biến
	LastName  string
	Age       int
	Address1  Address  //value
	Address2  *Address //pointer
}

type Address struct {
	Country string
	City    string
}

//Truyền vào con trỏ tham chiếu
func (p *Person) FullName() string {
	fmt.Printf("pointer receiver %p\n", p) //In ra địa chỉ trong vùng nhớ của con trỏ p
	p.Age = 100
	return p.FirstName + " " + p.LastName
}

//Truyền vào biến
func (p Person) String() string {
	fmt.Printf("value receiver : %p\n", &p) //In ra địa chỉ trong vùng nhớ của biến p
	p.Age = 200
	return fmt.Sprintf("%v is %v years old", p.FirstName, p.Age)
}

func NewAPerson(firstName string, lastName string, age int) *Person {
	if age < 0 {
		return nil
	}
	p := new(Person)
	p.FirstName = firstName
	p.LastName = lastName
	p.Age = age
	return p
}

//-----Fluent API
func BuildPerson() *Person {
	return new(Person)
}

func (p *Person) WithFirstName(firstName string) *Person {
	p.FirstName = firstName
	return p
}

func (p *Person) WithLastName(lastName string) *Person {
	p.LastName = lastName
	return p
}

func (p *Person) WithAge(age int) *Person {
	p.Age = age
	return p
}

func demoStruct() {
	person := Person{"Trinh", "Cuong", 45,
		Address{"Vietnam", "Hanoi"},
		&Address{"USA", "California"}}

	fmt.Printf("Địa chỉ con trỏ ban đầu %p\n", &person)

	fmt.Println(person.FullName())

	//fmt.Println(person.Age)

	fmt.Println(person.String())
	//fmt.Println(person.Age)

	//Sử dụng constructor. Trong Golang có cú pháp constructor chuẩn
	//mà lập trình viên tự định nghĩa constructor
	//Chúng ta có thể đặt tên hàm constructor là NewPerson hay NewAPerson
	//Ngược lại trong Java tên constructor luôn phải trùng tên với Java constructor
	/*	tom := NewAPerson("Trinh", "Cuong", -2)
		if tom != nil {
			fmt.Println(tom.FullName())
		}

		//Fluent API bản chất chỉ là cách viết nối tiếp các method mà thôi
		jack := BuildPerson().WithFirstName("Jack").WithLastName("London").WithAge(12)
		fmt.Println(jack)*/
}
