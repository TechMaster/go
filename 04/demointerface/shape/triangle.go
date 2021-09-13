package shape

import (
	"fmt"
	"math"
)

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := 0.5 * t.Perimeter()
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle A=%.2f, B=%.2f, C=%.2f", t.A, t.B, t.C)
}
