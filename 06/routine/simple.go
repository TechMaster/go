package main

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/goroutines
*/
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func demoGoRoutineSimple() {
	f("direct") //Gọi trực tiếp thực thi tuần tự

	go f("goroutine")

	go func(msg string) {
		// time.Sleep(8 * time.Millisecond) Hãy thử thêm lệnh này vào
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
