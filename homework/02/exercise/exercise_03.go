package exercise

// Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
// Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]


// Cách 1 sử dụng map
func RemoveDuplicates_01(numbers []int) (result []int) {
	keys := make(map[int]bool)
 
    for _, entry := range numbers {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            result = append(result, entry)
        }
    }
    return result
}

// Cách 2 : Duyệt qua từng phần tử của mảng
// Kiểm tra phần tử đó có nằm trong mảng result hay không
// Nếu không có thì append vào
func RemoveDuplicates_02(numbers []int) (result []int) {
	for _, number := range numbers {
		if !contains(result, number) {
			result = append(result, number)
		}
	}

	return result
}

// Function kiểm tra xem một phần tử có nằm trong mảng hay không
func contains(numbers []int, element int) bool {
	for _, number := range numbers {
		if number == element {
			return true
		}
	}
	return false
}