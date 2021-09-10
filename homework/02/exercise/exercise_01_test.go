package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var numbers = []int{1,2,3,5,2,5,4}

// Unit test
func Test_Max2Numbers_01(t *testing.T) {
	assert := assert.New(t)

	actualResult := Max2Numbers_01(numbers)
	desiredResult := 4

	assert.Equal(actualResult, desiredResult)
}

func Test_Max2Numbers_02(t *testing.T) {
	assert := assert.New(t)

	actualResult := Max2Numbers_02(numbers)
	desiredResult := 4

	assert.Equal(actualResult, desiredResult)
}

func Test_Max2Numbers_03(t *testing.T) {
	assert := assert.New(t)

	actualResult := Max2Numbers_03(numbers)
	desiredResult := 4

	assert.Equal(actualResult, desiredResult)
}

func Test_Max2Numbers_04(t *testing.T) {
	assert := assert.New(t)

	actualResult := Max2Numbers_04(numbers)
	desiredResult := 4

	assert.Equal(actualResult, desiredResult)
}

// Benchmark

func Benchmark_Max2Numbers_01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = Max2Numbers_01(numbers)
	}
}

func Benchmark_Max2Numbers_02(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = Max2Numbers_02(numbers)
	}
}

func Benchmark_Max2Numbers_03(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = Max2Numbers_03(numbers)
	}
}

func Benchmark_Max2Numbers_04(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = Max2Numbers_04(numbers)
	}
}