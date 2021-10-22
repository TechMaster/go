package main

import (
	"demoerror/foo"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	err := foo.Foo()
	fmt.Println(err)
	//demopanic.Connect_to_db()

	time.Sleep(10 * time.Second)
	fmt.Println("I am still alive")
}
