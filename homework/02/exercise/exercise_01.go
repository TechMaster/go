package exercise

import (
	"sort"
)

// Bài 1 Viết function tìm ra số lớn thứ nhì trong mảng các số.
// Ví dụ: max2Numbers([2, 1, 3, 4]) => 3

// Cách 1: Sắp xếp theo chiều tăng dần rồi lấy phần tử nhỏ thứ 2
func Max2Numbers_01(numbers []int) (result int) {
	// Sắp xếp mảng giảm dần
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})

	// Lặp qua mảng trả về phần tử đầu tiên khác phần tử lớn nhất
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[0] {
			result = numbers[i]
			break
		}
	}

	return result
}


// Cách 2
// Tìm giá trị lớn nhất trong mảng => loại các số có giá trị lớn nhất ra khỏi mảng
// Tìm lại số lớn nhất trong mảng lần nữa
func findMaxNumberArr(numbers []int) int {
	var max int = numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func Max2Numbers_02(numbers []int) int {
	var max = findMaxNumberArr(numbers)

	var newSlice []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != max {
			newSlice = append(newSlice, numbers[i])
		}
	}

	return findMaxNumberArr(newSlice)
}

// Cách 3
func Max2Numbers_03(numbers []int) int {
	// Remove các giá trị duplicates
	numbersRemoveDuplicates := RemoveDuplicates_01(numbers)

	// Sắp xếp mảng tăng dần
	sort.Ints(numbersRemoveDuplicates)

	// Lấy giá trị thứ 2 từ cuối lên
	return numbersRemoveDuplicates[len(numbersRemoveDuplicates) - 2]
}

// Cách 4
func Max2Numbers_04(numbers []int) int {
	var max = findMaxNumberArr(numbers)

	max2 := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max2 && numbers[i] < max {
			max2 = numbers[i]
		}
	}

	return max2
}