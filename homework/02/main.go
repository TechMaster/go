package main

import (
	"fmt"
	"homework-02/exercise"
)


func main() {
	// Bài 1
	numbers := []int{1,2,5,4,5,4,1}
	fmt.Println(exercise.Max2Numbers_01(numbers))
	fmt.Println(exercise.Max2Numbers_02(numbers))
	fmt.Println(exercise.Max2Numbers_03(numbers))
	fmt.Println(exercise.Max2Numbers_04(numbers))

	// Bài 2
	arrString := []string{"aba", "aa", "ad", "c", "vcd"}
	fmt.Println(exercise.FindMaxLengthElement_01(arrString))
	fmt.Println(exercise.FindMaxLengthElement_02(arrString))

	// Bài 3
	fmt.Println(exercise.RemoveDuplicates_01(numbers))
	fmt.Println(exercise.RemoveDuplicates_02(numbers))

	// Bài 4
	employees := []exercise.Employee{
		{"Hang", 5, 1000000},
		{"An", 7, 1500000},
		{"Cong", 10, 1300000},
		{"Ngoc", 7, 2000000},
		{"Oanh", 8, 2500000},
		{"Tam", 8, 2500000},
	}

	// Sắp xếp theo tên tăng dần theo bảng chữ cái
	fmt.Println(exercise.SortByName(employees))

	// Sắp xếp nhân viên theo mức lương giảm dần
	fmt.Println(exercise.SortBySalary(employees))

	// Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
	fmt.Println(exercise.Max2Salary(employees))
}