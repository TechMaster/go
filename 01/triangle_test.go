package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTriangle(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsTriangle(3, 4, 5))
}

func TestNotIsTriangle(t *testing.T) {
	assert := assert.New(t)
	assert.False(IsTriangle(3, 4, 15))
}
