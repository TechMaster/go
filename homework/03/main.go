package main

import (
	"fmt"
	"hw03/calendar"
	"hw03/docker"
	"hw03/tree"
)

func main() {
	// Bài 1
	calendar.ShowCalendar(2021, 9)
	fmt.Println()

	// Bài 2
	containers, _ := docker.ListContainer(docker.ALL)
	for _, c := range containers {
		fmt.Printf("%s - %s\n", c.Names[0], c.State)
	}

	_ = docker.StartContainerById("2926490de30f")
	_ = docker.StopContainerById("7f86a07c6921")

	fmt.Println()

	// Bài 3
	data := tree.PrintDirectoryTree(".")
	fmt.Println(data)
}