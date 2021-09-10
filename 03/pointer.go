package main

import "fmt"

/*
Hàm này không thực sự tăng giá trị
*/
func incIntValue(a int) {
	a++
	fmt.Println("Inside func a = ", a)
}

/*
Hàm này mới tăng giá trị a
*/
func incIntPointer(a *int) {
	(*a)++
	fmt.Println("Inside func a = ", *a)
}

/*
Hoán đổi vị trí ký tự đầu tiên và ký tự cuối cùng
*/
func modifyString(s string) {
	//s[len(s)-1:]: phần tử cuối cùng của chuỗi s
	//s[1:len(s)-1]: lấy chuỗi từ ký tự thứ 2 (vị trí 1) đến ký tự trước ký tự cuối cùng
	//s[:1]: phần tử đầu tiên
	s = s[len(s)-1:] + s[1:len(s)-1] + s[:1]
	//fmt.Println("Inside func s = ", s)
}

/*
Truyền con trỏ đến string
*/
func modifyString2(s *string) {
	/*	temp := *s
	*s = temp[len(temp)-1:] + temp[1:len(temp)-1] + temp[:1]*/

	*s = (*s)[len(*s)-1:] + (*s)[1:len(*s)-1] + (*s)[:1]
	//fmt.Println("Inside func s = ", s)
}

/*
Trả về chuỗi sau khi thay đổi
*/
func modifyString3(s string) string {
	return s[len(s)-1:] + s[1:len(s)-1] + s[:1]
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
	new_s := modifyString3(s)
	fmt.Println(new_s)
	modifyString2(&s)
	fmt.Println("Outside func s = ", s)

	x := 10
	y := &x
	fmt.Println("y = ", y)
	fmt.Println("*y = ", *y)
}
