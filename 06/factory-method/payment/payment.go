package payment

import "fmt"

// Interface
type PaymentMethod interface {
	Pay(int)
}

// Concrete product
type CashPaymentMethod struct{}
type DebitCardPaymentMethod struct{}

func (c *CashPaymentMethod) Pay(amount int) {
	fmt.Println("Thanh toán tiền mặt - Số tiền : ", amount)
}

func (c *DebitCardPaymentMethod) Pay(amount int) {
	fmt.Println("Thanh toán bằng thẻ ghi nợ - Số tiền : ", amount)
}

// Factory
const (
	CashType = 1
	DebitCardType = 2 
)

func GetPaymentMethod(paymentType int) PaymentMethod {
	switch paymentType {
	case CashType:
	  return new(CashPaymentMethod)
	case DebitCardType:
	  return new(DebitCardPaymentMethod)
	}
	return nil
}

// Client
func FactoryPayment() {
	paymentMT := GetPaymentMethod(DebitCardType)
	paymentMT.Pay(200000)
	
	paymentMT = GetPaymentMethod(CashType)
	paymentMT.Pay(500000)
}