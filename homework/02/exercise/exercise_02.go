package exercise

import "sort"

// Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.
// Ví dụ: findMaxLengthElement["aba", "aa", "ad", "c", "vcd"] => ["aba", "vcd"]

// Cách 1
// Tìm độ dài lớn nhất
// Tìm phần tử có độ dài = độ dài lớn nhất

func findMaxLength(arr []string) (maxLength int) {
	maxLength = len(arr[0])
	for _, ele := range arr {
		if len(ele) > maxLength {
			maxLength = len(ele)
		}
	}
	return maxLength
}

func FindMaxLengthElement_01(arr []string) (result []string) {
	maxLength := findMaxLength(arr)

	for _, ele := range arr {
		if len(ele) == maxLength {
			result = append(result, ele)
		}
	}

	return result
}

// Cách 2: Sắp xếp mảng giảm dần theo độ dài của phần tử.
// Tìm phần tử có độ dài bằng độ dài của phần tử đầu tiên
func FindMaxLengthElement_02(arr []string) (result []string) {
	// Sắp xếp
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) > len(arr[j])
	})

	// Tìm phần tử
	for _, ele := range arr {
		if len(ele) == len(arr[0]) {
			result = append(result, ele)
		} else {
			break
		}
	}

	return result
}