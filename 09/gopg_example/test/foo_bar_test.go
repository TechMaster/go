package test

import (
	"fmt"
	"gopgdemo/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_foo_bar(t *testing.T) {
	foos, err := repo.GetFooBar()
	assert.Nil(t, err)
	for _, foo := range foos {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println("  " + bar.Name)
		}
	}
	assert.Positive(t, len(foos))
}

func Test_select_foo_by_id(t *testing.T) {
	foo, err := repo.GetFooByID("ox-01")
	assert.Nil(t, err)
	if err == nil {
		fmt.Println(foo.Name)
		for _, bar := range foo.Bars {
			fmt.Println("  " + bar.Name)
		}
	}
}
