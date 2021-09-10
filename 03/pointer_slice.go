package main

import "fmt"

func demoValueSlice(a []int) {
	a[0] = -1
	fmt.Printf("Address of a = %p\n", a)
}
func demoPointerSlice(a *[]int) {
	(*a)[0] = 100
	fmt.Printf("Address of a = %p\n", a)
}

func demoValueVsPointerSlice() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Address of a = %p\n", a)
	fmt.Println(a)
	demoValueSlice(a)

	fmt.Println("------")
	demoPointerSlice(&a)
	fmt.Println(a)
}
