package repo

import (
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
)

type Address struct {
	Line    string
	City    string
	Country string
}

type Person struct {
	Name          string
	Emails        []string
	Date_of_birth time.Time
	Addresses     []*fake.AddressInfo
}

type Book struct {
	Title  string
	Author string
	Pages  int
}
