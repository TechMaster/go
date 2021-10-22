package main

import (
	"demoeris/bar"
	"demoeris/foo"
	"fmt"

	"github.com/rotisserie/eris"
)

func simple_print() {
	if err := foo.Foo(); err != nil {
		fmt.Println(err)
	}
}

func print_stack_trace() {
	if err := foo.Foo(); err != nil {
		format := eris.NewDefaultStringFormat(eris.FormatOptions{
			InvertOutput: true, // flag that inverts the error output (wrap errors shown first)
			WithTrace:    true, // flag that enables stack trace output
			InvertTrace:  true, // flag that inverts the stack trace output (top of call stack shown first)
		})

		// format the error to a string and print it
		formattedStr := eris.ToCustomString(err, format)
		fmt.Println(formattedStr)
	}
}

func check_error_type() {
	connect_api_failed := eris.New("connect API failed")
	err := bar.Call_other_rest()
	if eris.Is(err, connect_api_failed) {
		fmt.Println(err)
	}
}
func main() {
	check_error_type()
}
