package main

import "fmt"

func (p Person) String() string {
	return fmt.Sprintf("{%s : %s : %s} lives at {%s, %s, %s}",
		p.Id, p.FullName, p.Email, p.Addr.Location, p.Addr.City, p.Addr.Country)

}
