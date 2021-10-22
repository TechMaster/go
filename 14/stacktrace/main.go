package main

import (
	"fmt"
	"stacktrace/error_stack"
)

func main() {
	err := error_stack.AppendStackTrace(testOne())
	fmt.Printf("%s", err)

	//t, _ := err.StackTrace()
	//fmt.Printf("\n\n%s", t)
}

//Lỗi ngoài cùng
func testOne() error {
	return fmt.Errorf("testOne: %w", error_stack.AppendStackTrace(testTwo()))

}

func testTwo() error {
	return fmt.Errorf("testTwo: %w", error_stack.AppendStackTrace(testThree()))
}

//Lỗi trong cùng
func testThree() error {
	return error_stack.AppendStackTrace(fmt.Errorf("internal error"))
}
