package person

import "fmt"

type Address struct {
	City string
	Stress string
}

type Person struct {
	Name string
	Email string
	Age int
	Address *Address
}

func(p *Person) Clone() iPerson{
	return &Person{
		Name: p.Name,
		Email: p.Email,
		Age: p.Age,
		Address: p.Address,
	}
}

func(p *Person) ChangeAddress(stress string) {
	p.Address.Stress = stress
}

func(p *Person) Print() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Email: %s\n", p.Email)
	fmt.Printf("Address: %v\n", p.Address)
	fmt.Println()
}