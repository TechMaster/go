package main

import (
	"demoerror/demopanic"
	"demoerror/foo"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	err := foo.Foo()
	fmt.Println(err)
	demopanic.Connect_to_db()
}
