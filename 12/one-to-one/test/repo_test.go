package test

import (
	"fmt"
	"one-to-one/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllPerson(t *testing.T) {
	db := repo.GetInstance()
	persons, err := repo.GetAllPerson(db)
	assert.Nil(t, err)

	for i, p := range persons {
		fmt.Printf("%d - Name : %s\n", i, p.FullName)
		fmt.Println("  Seri : " + p.Card.Seri)
	}
	assert.Positive(t, len(persons))
}

func Test_GetPersonById(t *testing.T) {
	db := repo.GetInstance()
	person, err := repo.GetPersonById(db, "1")
	assert.Nil(t, err)

	fmt.Println("Name : " + person.FullName)
	fmt.Println("  Seri : " + person.Card.Seri)
}

func Test_CreateData(t *testing.T) {
	db := repo.GetInstance()
	err := repo.CreateData(db)
	assert.Nil(t, err)
}
