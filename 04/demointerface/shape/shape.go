package shape

import "fmt"

type Shape interface {
	Perimeter() float64
	Area() float64
	String() string
}

func DemoShape() {
	rec := Rectangle{
		Width:  10,
		Height: 20,
	}

	var shape Shape
	shape = rec
	fmt.Println(shape.Area())

	cir := Circle{
		Radius: 10,
	}

	tri := Triangle{
		A: 3,
		B: 4,
		C: 5,
	}

	shapes := []Shape{rec, cir, tri}
	for _, shape := range shapes {
		//Polymorphism
		fmt.Printf("%s perimeter = %.2f\n", shape, shape.Perimeter())
	}
}
