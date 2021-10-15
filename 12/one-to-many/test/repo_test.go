package test

import (
	"fmt"
	"one-many/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFooBar(t *testing.T) {
	db := repo.GetInstance()
	foos, err := repo.GetFooBar(db)
	assert.Nil(t, err)
	for _, foo := range foos {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println("  " + bar.Name)
		}
	}
	assert.Positive(t, len(foos))
}

func Test_GetFooById(t *testing.T) {
	db := repo.GetInstance()
	foo, err := repo.GetFooById(db, "ox-01")
	assert.Nil(t, err)
	fmt.Println(foo.Name)
	for _, bar := range foo.Bars {
		fmt.Println("  " + bar.Name)
	}
}

func Test_GetBarById(t *testing.T) {
	db := repo.GetInstance()
	bar, err := repo.GetBarById(db, "bar1")
	assert.Nil(t, err)

	fmt.Println(bar.FooId)
	fmt.Println("  " + bar.Id)
	fmt.Println("  " + bar.Name)
}

func Test_CreateData(t *testing.T) {
	db := repo.GetInstance()
	err := repo.CreateData(db)
	assert.Nil(t, err)
}
