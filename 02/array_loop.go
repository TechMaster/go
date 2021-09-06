package main

import "fmt"

func indexLoop() {
	cars := []string{"Toyota", "Mercedes", "BMW"}
	for i := 0; i < len(cars); i++ {
		fmt.Println(i, cars[i])
	}
}

/*
Golang không có while. Less is Better
*/
func whileLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	i := 0
	for i < len(cars) {
		fmt.Println(i, cars[i])
		i++
	}
}

func rangeLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	//fmt.Println(cars[0]) // Toyota
	for index, car := range cars {
		fmt.Println(index, car)
	}
}

/*
Sử dụng defer khi hàm thoát thì gỡ dần ngăn sếp (stack) để thực hiện
Chỉ áp dụng với số phần tử mảng nhỏ.
Cách này hơi hack não
*/
func reverseLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	//fmt.Println(cars[0]) // Toyota
	for index, car := range cars {
		defer fmt.Println(index, car)
	}
}

/*
Tốc độ thực thi nhanh nhất
*/
func rawReverseLoop() {
	cars := []string{"Toyota", "Mercedes", "BMW"}
	for i := len(cars) - 1; i >= 0; i-- {
		fmt.Println(i, cars[i])
	}
}
