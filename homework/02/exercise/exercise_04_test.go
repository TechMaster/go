package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var employees = []Employee{
	{"Hang", 5, 1000000},
	{"An", 7, 1500000},
	{"Cong", 10, 1300000},
	{"Ngoc", 7, 2000000},
	{"Oanh", 8, 2500000},
	{"Tam", 8, 2500000},
}

// Unit test
func Test_SortByName(t *testing.T) {
	assert := assert.New(t)

	actualResult := SortByName(employees)
	desiredResult := []Employee{
		{"An", 7, 1500000},
		{"Cong", 10, 1300000},
		{"Hang", 5, 1000000},
		{"Ngoc", 7, 2000000},
		{"Oanh", 8, 2500000},
		{"Tam", 8, 2500000},
	}

	assert.Equal(actualResult, desiredResult)
}

func Test_SortBySalary(t *testing.T) {
	assert := assert.New(t)

	actualResult := SortBySalary(employees)
	desiredResult := []Employee{
		{"Cong", 10, 1300000},
		{"Oanh", 8, 2500000},
		{"Tam", 8, 2500000},
		{"Ngoc", 7, 2000000},
		{"An", 7, 1500000},
		{"Hang", 5, 1000000},
	}

	assert.Equal(actualResult, desiredResult)
}

func Test_Max2Salary(t *testing.T) {
	assert := assert.New(t)

	actualResult := Max2Salary(employees)
	desiredResult := []Employee{
		{"Oanh", 8, 2500000},
		{"Tam", 8, 2500000},
	}

	assert.Equal(actualResult, desiredResult)
}