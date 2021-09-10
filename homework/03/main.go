package main

import (
	"hw03/calendar"
	"hw03/docker"
)

func main() {
	calendar.ShowCalendar(2021, 9)
	docker.ListContainer()
}
