package employee

import "fmt"

func ClientBuildEmployee() {
	direc := Director{}

	// Build sale
	sale := &Sale{}
	direc.SetBuilder(sale)
	employee := direc.BuildEmployee()

	// In thông tin sale
	printDetail(employee)

	// Build teacher
	teacher := &Teacher{}
	direc.SetBuilder(teacher)
	employee = direc.BuildEmployee()

	// In thông tin teacher
	printDetail(employee)
}

func printDetail(e Employee) {
	fmt.Println("Name: ", e.Name)
	fmt.Println("Email: ", e.Email)
	fmt.Println("Age: ", e.Age)
	fmt.Println("Position: ", e.Position)
}