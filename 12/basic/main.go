package main

import (
	"basic/repo"
)

func main() {
	// Connect database
	repo.ConnectDB()

	// repo.DemoCreate()
	// repo.DemoSelect()
	// repo.DemoUpdate()
	// repo.DemoDelete()
	repo.DemoRawQuery()
}
