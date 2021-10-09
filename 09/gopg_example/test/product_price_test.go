package test

import (
	"gopgdemo/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_product_price(t *testing.T) {
	err := repo.Create_product_price()
	assert.Nil(t, err)
}
