package basic

import "fmt"

func DemoInterfaceTypeAssert() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/* Ví dụ này sẽ gây lỗi
	f = i.(float64) // panic
	fmt.Println(f)
	*/

	//Kiểm tra đúng kiểu rồi mới in
	if f, ok := i.(string); ok { //Golang if assignment statement
		fmt.Println(f)
	}

	person := Person{
		Id:   "OX-13",
		Name: "Donald Trump",
	}

	i = person

	if f, ok := i.(Person); ok {
		fmt.Println(f.Id) //Hãy đặt break point ở đây
	}

}

func DemoSwitchType() {
	person := Person{
		Id:   "OX-13",
		Name: "Donald Trump",
	}

	arrayAnyType := []interface{}{
		3.14, "Hello", true, person, Add, []string{"John", "Anna", "Tom"},
	}

	//-----
	fmt.Println("--------------------------------")
	for _, v := range arrayAnyType {
		switch v.(type) {
		case string:
			fmt.Printf("%v is string\n", v)
		case Person:
			fmt.Printf("%v is Person\n", v)
		case func(int, int) int:
			fmt.Printf("%v is func(int, int) int\n", v)
		}

	}
}
