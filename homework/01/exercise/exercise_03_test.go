package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var number = 20
var desiredResult = []int{2, 3, 5, 7, 11, 13, 17, 19}

// Unit test
func Test_FindPrimes_01(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindPrimes_01(number)

	assert.Equal(actualResult, desiredResult)
}

func Test_FindPrimes_02(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindPrimes_02(number)

	assert.Equal(actualResult, desiredResult)
}

func Test_FindPrimes_03(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindPrimes_03(number)

	assert.Equal(actualResult, desiredResult)
}

func Test_FindPrimes_04(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindPrimes_04(number)

	assert.Equal(actualResult, desiredResult)
}

func Test_FindPrimes_05(t *testing.T) {
	assert := assert.New(t)

	actualResult := FindPrimes_05(number)

	assert.Equal(actualResult, desiredResult)
}

// Benchmark

func Benchmark_FindPrimes_01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindPrimes_01(number)
	}
}

func Benchmark_FindPrimes_02(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindPrimes_02(number)
	}
}

func Benchmark_FindPrimes_03(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindPrimes_03(number)
	}
}

func Benchmark_FindPrimes_04(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindPrimes_04(number)
	}
}

func Benchmark_FindPrimes_05(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = FindPrimes_05(number)
	}
}