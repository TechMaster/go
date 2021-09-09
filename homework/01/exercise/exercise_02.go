package exercise

import (
	"fmt"
	"math/rand"
)

/*
Bài 2: Đoán số
Máy tính tự sinh ra một số nguyên dương X >= 0 và <= 100. Lập trình một vòng lặp để người dùng đoán số.

- Nếu số đoán lớn hơn X thì in ra "Số bạn đoán lớn hơn X"
- Nếu số đoán nhỏ hơn X thì in ra "Số bạn đoán nhỏ hơn X"
- Nếu bằng X thì in ra "Bạn đã đoán đúng"
*/

// Hàm so sánh 2 số
func CompareNumber(computerNumber, userNumber int) {
	if computerNumber < userNumber {
		fmt.Printf("Số bạn đoán lớn hơn %d\n", computerNumber)
	} else if computerNumber > userNumber {
		fmt.Printf("Số bạn đoán nhỏ hơn %d\n", computerNumber)
	} else {
		fmt.Println("Bạn đã đoán đúng")	
	}
}

// Random 1 số trong khoảng min -> max
func RandomNumber(min, max int) int {
	return rand.Intn(max - min + 1) + min
}