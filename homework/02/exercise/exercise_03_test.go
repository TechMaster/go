package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var test = []int{1,2,5,2,6,2,5}
var desired = []int{1,2,5,6}

// Unit test
func Test_RemoveDuplicates_01(t *testing.T) {
	assert := assert.New(t)

	actualResult := RemoveDuplicates_01(test)

	assert.Equal(actualResult, desired)
}

func Test_RemoveDuplicates_02(t *testing.T) {
	assert := assert.New(t)

	actualResult := RemoveDuplicates_02(test)

	assert.Equal(actualResult, desired)
}

// Benchmark

func Benchmark_RemoveDuplicates_01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = RemoveDuplicates_01(test)
	}
}

func Benchmark_RemoveDuplicates_02(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = RemoveDuplicates_02(test)
	}
}