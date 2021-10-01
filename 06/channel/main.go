package main

import "fmt"

func main() {
	DemoChannelString()
}

func DemoChannelString() {
	pipe := make(chan string, 2)
	go func() {
		pipe <- "water"
		pipe <- "beer"
		pipe <- "alcohol"
		pipe <- "tea"
		close(pipe)
	}()

	for receiver := range pipe {
		fmt.Println(receiver)
	}
}
