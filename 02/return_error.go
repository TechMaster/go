package main

import (
	"fmt"
	"math"
)

/*
Ví dụ trả về lỗi
*/
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math: square root of negative number %g", f)
	}
	return math.Sqrt(f), nil
}
func demoError() {
	if result, err := Sqrt(-1); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
