package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  Address //value
}

type Address struct {
	Country string
	City    string
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

//-----Fluent API--------
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
		Address{"Vietnam", "Hanoi"}}

	fmt.Println(person)

	//Fluent API bản chất chỉ là cách viết nối tiếp các method mà thôi
	jack := BuildPerson().WithFirstName("Jack").WithLastName("London").WithAge(12)
	fmt.Println(jack)
}
