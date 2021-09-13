package basic

import (
	"fmt"
)

/*
type Stringer interface {
    String() string
}
*/
type Operator2operand = func(int, int) int
type AnyType = interface{}

type Product struct {
	Id    string
	Name  string
	Price int
}

func (p Product) String() string {
	return fmt.Sprintf("Id : %v - Name : %v - Price : %v", p.Id, p.Name, p.Price)
}

type Point struct{ X, Y float64 }
type Color struct{ R, G, B int }
type ColoredPoint struct {
	// thuộc tính ẩn danh
	Point

	// thuộc tính bình thường
	Color Color
}

func (cp ColoredPoint) String() string {
	return fmt.Sprintf("Point at %.2f - %.2f has color {%d, %d, %d}", cp.Point.X, cp.Point.Y, cp.Color.R, cp.Color.G, cp.Color.B)

	/* return fmt.Sprintf("Point at %.2f - %.2f has color {%d, %d, %d}", cp.X, cp.Y, cp.Color.R, cp.Color.G, cp.Color.B)
	sử dụng cú pháp truy xuất thuộc tính ẩn danh cho kết quả tương tự */
}

/*
Một số kiểu qua căn bản có sẵn hàm String() rồi nên không thể định nghĩa lái
invalid receiver int (basic or unnamed type)

func (p int) String() string {
	return "it is basic int type"
}

func (array []interface{}) String() string {

}
*/

func DemoStringerInterface() {
	product := Product{
		Id:    "ax123",
		Name:  "Sony Walkman",
		Price: 450000,
	}
	fmt.Println(product)

	//Không cần định nghĩa lại hàm string
	arrayAnyType := []interface{}{
		3.14, "Hello", true, product, Add, []string{"John", "Anna", "Tom"},
	}

	fmt.Println(arrayAnyType)

	cp := ColoredPoint{
		Point: Point{X: 10.5, Y: 20.1},
		Color: Color{100, 255, 255},
	}

	fmt.Println(cp)

}
