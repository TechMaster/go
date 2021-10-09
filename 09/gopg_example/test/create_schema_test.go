package test

import (
	"gopgdemo/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_create_schema(t *testing.T) {
	err := repo.CreateSchema()
	assert.Nil(t, err)
}
