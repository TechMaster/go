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
func testOne() (err error) {
	err = testTwo()
	if err != nil {
		return fmt.Errorf("testOne: %w", error_stack.AppendStackTrace(err))
	} else {
		return nil
	}

}

func testTwo() (err error) {
	err = testThree()
	if err != nil {
		return fmt.Errorf("testTwo: %w", error_stack.AppendStackTrace(err))
	} else {
		return nil
	}
}

//Lỗi trong cùng
func testThree() error {
	return error_stack.AppendStackTrace(fmt.Errorf("internal error"))
}
