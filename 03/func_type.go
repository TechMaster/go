package main

import "fmt"

//Định nghĩa kiểu hàm
type Operator func(a int, b int) int

type POperator *Operator

func Add(a int, b int) int {
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}

//Dùng Value receiver rất đơn giản
func (op Operator) compute(a int, b int) int {
	return op(a, b)
}

//Dùng Pointer receiver phức tạp hơn
func (op *Operator) compute2(a int, b int) int {
	return (*op)(a, b)
}

func DemoOperator() {
	var op Operator
	op = func(a int, b int) int {
		return a + b
	}

	fmt.Println(op.compute(1, 2))
	fmt.Println(op.compute2(1, 2))

	op = Subtract
	fmt.Println(op.compute(10, 5))
	fmt.Println(op.compute2(10, 5))

	var op2 POperator
	op2 = &op
	fmt.Println((*op2).compute(10, 5))
}
