package exercise_02

type Contract struct {
    empId  int
    basicpay int
}

func (c *Contract) GetSalary() int {
	return c.basicpay 
}

func (c *Contract) ChangeEmId(empId int) {
	c.empId = empId
}

func (c *Contract) ChangeBasicPay(basicPay int) {
	c.basicpay = basicPay
}

func (c *Contract) ChangePf(pf int) {}

func (c *Contract) Clone() Employee {
	return &Contract{
		empId: c.empId,
		basicpay: c.basicpay,
	}
}