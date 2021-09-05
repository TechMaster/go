package main

import "fmt"

func demoSlice() {
	letters := []string{"a", "b", "c", "d"}
	letters = append(letters, "e")
	length := len(letters)
	fmt.Println(length)
	fmt.Println(letters[:2])
}
