package main

import "fmt"

func incIntValue(a int) {
	a++
	fmt.Println("Inside func a = ", a)
}
func incIntPointer(a *int) {
	(*a)++
	fmt.Println("Inside func a = ", *a)
}

func modifyString(s string) {
	s = s[len(s)-1:] + s[1:len(s)-1] + s[:1]
	fmt.Println("Inside func s = ", s)
}

func DemoPointer() {
	a := 10
	incIntValue(a)
	fmt.Println("Outside func a = ", a)

	incIntPointer(&a)
	fmt.Println("Outside func a = ", a)

	b := new(int)
	fmt.Printf("b address %p\n", b)
	*b = 10
	incIntPointer(b)
	fmt.Println("Outside func b = ", *b)

	var c *int
	fmt.Printf("c address %p\n", c)

	s := "hello"
	modifyString(s)
	fmt.Println("Outside func s = ", s)

	x := 10
	y := &x
	fmt.Println("y = ", y)
	fmt.Println("*y = ", *y)
}
