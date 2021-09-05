package main

import "fmt"

func demoArray2D() {
	langs := [][]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "RUST", "Crystal", "OCAML"}}
	for _, v := range langs {
		for _, lang := range v {
			fmt.Print(lang, " ")
		}
		fmt.Println()
	}
}
