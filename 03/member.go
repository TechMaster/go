package main

type Member struct {
	id      string
	name    string
	email   string
	pass    string
	roles   []string
	age     int
	enabled bool
}

func PassStructAsValue(p Member) {
	p.id = "100"
	p.email = "cuong@techmaster.vn"
	p.pass = "2221432sdv"
	p.name = "John"
	p.roles = []string{"admin", "trainer"}
	p.age = 50
	p.enabled = true
}

func PassStructAsPointer(p *Member) {
	p.id = "100"
	p.email = "cuong@techmaster.vn"
	p.pass = "2221432sdv"
	p.name = "John"
	p.roles = []string{"admin", "trainer"}
	p.age = 50
	p.enabled = true
}
