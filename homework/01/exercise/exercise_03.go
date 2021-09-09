package exercise

import (
	"math"
	"math/big"
)

/*
Bài 3: Lập dãy số nguyên tố
Nhập vào số nguyên dương N < 100,000 hãy trả về mảng các số nguyên tố
*/

// Cách 1
func isPrime_01(num int) bool {

	// Số nguyên tố bắt đầu từ 2 trở đi
	if num < 2 {
		return false
	}

	// duyệt từ 2 -> sqrt(num) => nếu chia hết cho số nào => false
	for i := 2; i <= int(math.Sqrt(float64(num))) ; i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

// Cách 2
func isPrime_02(num int) bool {
	if num < 2 {
		return false
	}

	// Số lượng ước ban đầu (khác 1 và chính nó)
	factors := 0

	for i := 2; i <= int(math.Sqrt(float64(num))) ; i++ {
		if num % i == 0 {
			factors++ // Nếu chia hết thì tăng số lượng ước lên 1
		}
	}

	return factors == 0
}

// Cách 3
func isPrime_03(num int) bool {
	if num < 2 {
		return false
	}

	// Kiểm soát số vòng lặp từ 2 -> sqrt(num)
	for i := 2; i*i <= num; {
		if num % i == 0 {
			return false
		}
        i += 1
	}
	return true
}

// Cách 4
func isPrime_04(num int) bool {
	if num < 2 {
		return false
	}

	// Đặt flag
	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num))) ; i++ {
		if num % i == 0{
			isPrime = false
			break
		}
	}

	return isPrime
}

// Tìm kiếm tất cả các số nguyến tố <= num (cách 1)
func FindPrimes_01(num int) (result []int) {
	for i := 2; i <= num; i++ {
		if isPrime_01(i) {
			result = append(result, i)
		}
	}
	return result
}

// Tìm kiếm tất cả các số nguyến tố <= num (cách 2)
func FindPrimes_02(num int) (result []int) {
	for i := 2; i <= num; i++ {
		if isPrime_02(i) {
			result = append(result, i)
		}
	}
	return result
}

// Tìm kiếm tất cả các số nguyến tố <= num (cách 3)
func FindPrimes_03(num int) (result []int) {
	for i := 2; i <= num; i++ {
		if isPrime_03(i) {
			result = append(result, i)
		}
	}
	return result
}

// Tìm kiếm tất cả các số nguyến tố <= num (cách 4)
func FindPrimes_04(num int) (result []int) {
	for i := 2; i <= num; i++ {
		if isPrime_04(i) {
			result = append(result, i)
		}
	}
	return result
}

// Tìm kiếm tất cả các số nguyến tố <= num (cách 5)
func FindPrimes_05(num int) (result []int) {
	for i := 2; i <= num; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) { // Sử dụng build-in method (Đầu vào < 2^64)
			result = append(result, i)
		}
	}
	return result
}