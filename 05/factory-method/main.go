package main

import (
	p "factory-pattern/post_factory"
	"factory-pattern/payment"
)

func main() {
	p.PostFactory()

	payment.FactoryPayment()
}
