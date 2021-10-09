package test

import (
	"fmt"
	"one_many/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_get_foo_bar(t *testing.T) {
	foos, err := repo.Get_Foo_Bar()
	assert.Nil(t, err)
	for _, foo := range foos {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println("  " + bar.Name)
		}
	}
	assert.Positive(t, len(foos))
}

func Test_get_foo_by_id(t *testing.T) {
	foo, err := repo.Get_Foo_By_ID("ox-01")
	assert.Nil(t, err)
	if err == nil {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println("  " + bar.Name)
		}
	}
}

func Test_get_bar_by_id(t *testing.T) {
	bar, err := repo.Get_Bar_By_ID("bar3")
	assert.Nil(t, err)
	if err == nil && bar != nil {
		fmt.Println(bar.Name)
		fmt.Println(bar.Foo.Name)
	}
}

func Test_create_foo_bar(t *testing.T) {
	err := repo.Create_Foo_Bar()
	assert.Nil(t, err)
}
