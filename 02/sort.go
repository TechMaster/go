package main

import (
	"fmt"
	"sort"
)

/*
Sắp xếp slice. Có 3 hàm sắp xếp chuẩn
sort.Ints
sort.Float64s
sort.Strings
*/
func sortIntSlice() {
	intSlice := []int{10, 5, 25, 351, 14, 9} // unsorted
	fmt.Println("Slice of integer BEFORE sort:", intSlice)
	sort.Ints(intSlice)
	fmt.Println("Slice of integer AFTER  sort:", intSlice)
}

/*
Sắp xếp thứ tự ngược
*/
func sortReverse() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

/*
Sort with custom function less
func Slice(x interface{}, less func(i, j int) bool)
Xem thêm https://pkg.go.dev/sort#Slice
*/
func sortSliceWithFunc() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}

/*
SliceStable sắp xếp slice sử dụng hàm less, những phần tử giống nhau được giữa nguyên thứ tự
*/
func sortSliceStable() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75}, //--
		{"Bob", 75},
		{"Alice", 75}, //--
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// Sort by age preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age,name:", people)
}

/*
Tìm kiếm int trong một mảng đã sắp xếp từ thấp lên cao
*/
func search() {
	a := []int{1, 2, 3, 4, 6, 7, 8}

	x := 2
	i := sort.SearchInts(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i, a)

	x = 5
	i = sort.SearchInts(a, x)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)
}
