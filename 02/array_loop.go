package main

import "fmt"

func indexLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
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

func reverseLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}
	//fmt.Println(cars[0]) // Toyota
	for index, car := range cars {
		defer fmt.Println(index, car)
	}
}
