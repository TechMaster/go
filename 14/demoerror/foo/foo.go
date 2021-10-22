package foo

import (
	"demoerror/bar"
	"fmt"
)

func Foo() error {
	fmt.Println("Foo function")
	return bar.Bar()
}
