package test

import (
	"gopgdemo/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_writer_story(t *testing.T) {
	err := repo.Create_writer_story()
	assert.Nil(t, err)
}
