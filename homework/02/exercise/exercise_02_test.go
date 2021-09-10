package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arrString = []string{"aba", "aa", "ad", "c", "vcd"}
var desiredResult = []string{"aba", "vcd"}

// Unit test
func Test_FindMaxLengthElement_01(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindMaxLengthElement_01(arrString)

	assert.Equal(actualResult, desiredResult)
}

func Test_FindMaxLengthElement_02(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindMaxLengthElement_02(arrString)

	assert.Equal(actualResult, desiredResult)
}

// Benchmark

func Benchmark_FindMaxLengthElement_01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindMaxLengthElement_01(arrString)
	}
}

func Benchmark_FindMaxLengthElement_02(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindMaxLengthElement_02(arrString)
	}
}