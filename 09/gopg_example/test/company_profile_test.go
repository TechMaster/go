package test

import (
	"gopgdemo/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_company_profile(t *testing.T) {
	err := repo.Create_company_profile()
	assert.Nil(t, err)
}
