package main

import "fmt"

/*
Kiểm tra 3 cạnh có tạo nên một tam giác hay không
*/
func IsTriangle(a float64, b float64, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}
func main() {
	fmt.Println(IsTriangle(3, 4, 15))
}
