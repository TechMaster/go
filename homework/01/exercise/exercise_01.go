package exercise

import (
	"fmt"
	"math"
)

/*
Bài 1: Giải phương trình bậc 2
Nhập vào ba số a, b, c kiểu float64 hãy giải phương trình bậc 2.
*/

func SolveEquation2(a, b, c float64) {
	// Tính delta
	delta := math.Pow(b, 2) - (4 * a * c)

	// Các điều kiện của delta
	// delta > 0 : 2 nghiệm phân biệt
	// delta = 0 : nghiệm duy nhất
	// delta < 0 : 2 nghiệm phức
	
	if delta > 0 {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("Phương trình có 2 nghiệm phân biệt x1 = %f, x2 = %f\n", x1, x2)
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Printf("Phương trình có nghiệm duy nhất x1 = x2 = %f\n", x)
	} else {
		imaginary := math.Sqrt(math.Abs(delta))
		x1 := complex(-b/(2*a), imaginary/(2*a))
		x2 := complex(-b/(2*a), -imaginary/(2*a))
		fmt.Printf("Phương trình 2 nghiệm phức x1 = %v, x2 = %v\n", x1, x2)
	}
}
