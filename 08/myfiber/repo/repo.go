package repo

import fake "github.com/brianvoe/gofakeit/v6"

var (
	people []*Person
)

func init() {
	people = []*Person{}
	for i := 0; i < 30; i++ {
		people = append(people, &Person{
			Name:          fake.Name(),
			Emails:        []string{fake.Email(), fake.Email()},
			Date_of_birth: fake.Date(),
			Addresses: []*fake.AddressInfo{
				fake.Address(), fake.Address(),
			},
		})
	}
}
func ListPeople() []*Person {
	return people
}
