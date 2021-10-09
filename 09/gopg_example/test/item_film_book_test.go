package test

import (
	"gopgdemo/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_film_book(t *testing.T) {
	err := repo.Create_films_books()
	assert.Nil(t, err)
}
