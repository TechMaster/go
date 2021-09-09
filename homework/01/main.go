package main

import (
	"fmt"
	"homework-01/exercise"
	"math/rand"
	"time"
)

func main() {
	// Bài 1: Giải phương trình bậc 2
	fmt.Println("Bài 1")

	exercise.SolveEquation2(1,2,3)
	exercise.SolveEquation2(1,4,4)
	exercise.SolveEquation2(1,5,4)

	// Bài 2: Đoán số
	fmt.Printf("\nBài 2\n")

	// Random số
	rand.Seed(time.Now().UnixNano())
	computerNumber := exercise.RandomNumber(0, 100)
	userNumber := 23

	exercise.CompareNumber(computerNumber, userNumber)

	// Bài 3: Tìm dãy số nguyên tố
	fmt.Printf("\nBài 3\n")

	fmt.Println(exercise.FindPrimes_01(20))
	fmt.Println(exercise.FindPrimes_02(20))
	fmt.Println(exercise.FindPrimes_03(20))
	fmt.Println(exercise.FindPrimes_04(20))
	fmt.Println(exercise.FindPrimes_05(20))
}