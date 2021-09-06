package main

import (
	"fmt"
)

func demoSlice() {
	a := []string{"a", "b", "c", "d"}
	a = append(a, "e")
	fmt.Println(len(a))
	fmt.Println(a[:2])        //Lấy 2 phần tử đầu tiên
	fmt.Println(a[2:])        //Bỏ qua 2 phần tử đầu tiên
	fmt.Println(a[len(a)-2:]) //Lấy 2 phần tử cuối cùng
	fmt.Println(a[1:3])
}

/*
Xoá nhanh không bảo toàn thứ tự phần tử trong slice
Hãy debug để hiểu thuật toán
*/
func removeItemSliceNotKeepOrder(a []string, i int) []string {
	a[i] = a[len(a)-1]  // Copy last element to index i.
	a[len(a)-1] = ""    // Erase last element (write zero value).
	return a[:len(a)-1] // Truncate slice.
}

func removeItemSliceNotKeepOrder2(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

/*
Xoá phần tử bảo toàn thứ tự
*/
func removeItemSliceKeepOrder(a []string, i int) []string {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}

/*
Xoá phần tử bảo toàn thứ tự. Cách 2
*/
func removeItemSliceKeepOrder2(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

/*
Đảo một slice. Chú ý khi truyền vào một mảng/slice là truyền by reference

When a slice is passed to a function, even though it's passed by value,
the pointer variable will refer to the same underlying array.
Hence when a slice is passed to a function as parameter, changes made inside
the function are visible outside the function too.
*/
func reverseNotKeepInput(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

/*
Đảo slice nhưng giữ nguyên slice ban đầu
*/
func reverseKeepInput(a []string) (reversed []string) {
	for i := range a {
		n := a[len(a)-1-i]
		reversed = append(reversed, n)
	}
	return
}

/*
Cách này cũng được nhưng code dài dòng. Phong cách của Golang là phải ngắn gọn, chạy nhanh
tốn ít bộ nhớ. Mọi sự dài dòng đều quy ra tiền cả !
*/
func reverseKeepInput2(a []string) []string {
	reversed := []string{}
	for i := range a {
		n := a[len(a)-1-i]
		reversed = append(reversed, n)
	}
	return reversed
}

/*
Loại bỏ phần tử lập lại trong slice
*/
func removeDuplicate(a []string) (result []string) {
	keys := map[string]bool{} //hoặc có thể viết keys := make(map[string]bool) nhưng dài dòng
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return
}

/*
Đảo mảng chứa các phần tử có kiểu bất kỳ interface{}
*/
func reverseSliceAnyType(a []interface{}) (reversed []interface{}) {
	for i := range a {
		n := a[len(a)-1-i]
		reversed = append(reversed, n)
	}
	return
}
