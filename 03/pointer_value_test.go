package main

import (
	"lesson03/pointer"
	"lesson03/value"
	"testing"
)

var p = Member{
	id:      "OX13",
	name:    "Valya",
	email:   "valya@gmail.com",
	pass:    "2334234j,m as",
	roles:   []string{"admin", "manager"},
	age:     18,
	enabled: false,
}

var nike = Product{
	id:    "ox-13",
	name:  "Nike Air shoes",
	price: 100,
}

func BenchmarkPassStructAsValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassStructAsValue(p)
	}
}

func BenchmarkPassStructAsPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassStructAsPointer(&p)
	}
}

//--------
func BenchmarkValueReceiver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nike.incPrice1(20)
	}
}

func BenchmarkPointerReceiver(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nike.incPrice2(20)
	}
}

//-------
func BenchmarkValueAccountRepo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value.InitData()
	}
}

func BenchmarkPointerAccountRepo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointer.InitData()
	}
}

func BenchmarkValueFindById(b *testing.B) {
	value.InitData()
	for i := 0; i < b.N; i++ {
		value.InitData()
	}
}

func BenchmarkPointerFindById(b *testing.B) {
	pointer.InitData()
	for i := 0; i < b.N; i++ {
		pointer.InitData()
	}
}
