package basic

import "fmt"

type Person struct {
	Id   string
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("Id : %v - Name : %v", p.Id, p.Name)
}

type Operator func(a int, b int) int

func DemoAnyType() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = Person{
		Id:   "OX-13",
		Name: "Donald Trump",
	}
	describe(i) //struct

	describe(&i) //pointer to struct

	describe([]string{"a", "b", "c", "d", "e", "f"})
	intArray := [3]int{1, 2, 3}
	intSlice := []int{1, 2, 3}

	describe(intArray)
	describe(intSlice[2:])
	describe(Add)

	const pi = 3.14156
	describe(pi)

	describe(map[string]interface{}{
		"id":   "OX-13",
		"name": "Ronaldo",
	})

}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func Add(a int, b int) int {
	return a + b
}
