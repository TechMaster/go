package main

import "fmt"

type Product struct {
	id    string
	name  string
	price int
}

/* Nâng giá sản phẩm lên
price = price * (1 + percentage/100)
Đây là value receiver function
*/
func (product Product) increasePrice1(percentage int) {
	fmt.Printf("%p\n", &product)
	product.price = product.price * (100 + percentage) / 100
}

/* Nâng giá sản phẩm lên
price = price * (1 + percentage/100)
Đây là pointer receiver function
*/
func (product *Product) increasePrice2(percentage int) {
	fmt.Printf("%p\n", product)
	product.price = product.price * (100 + percentage) / 100
}

//Dùng để kiểm tra benchmark
func (product Product) incPrice1(percentage int) {
	product.price = product.price * (100 + percentage) / 100
}

func (product *Product) incPrice2(percentage int) {
	product.price = product.price * (100 + percentage) / 100
}

func demoProduct() {
	nike := Product{
		id:    "ox-13",
		name:  "Nike Air shoes",
		price: 100,
	}
	fmt.Printf("%p\n", &nike)

	nike.increasePrice1(20)
	fmt.Println(nike.price)

	nike.increasePrice2(20)
	fmt.Println(nike.price)
}
