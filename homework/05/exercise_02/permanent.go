package exercise_02


type Permanent struct {
    empId    int
    basicpay int
    pf       int
}

func (p *Permanent) GetSalary() int {
	return p.basicpay + p.pf
}

func (p *Permanent) ChangeEmId(empId int) {
	p.empId = empId
}

func (p *Permanent) ChangeBasicPay(basicPay int) {
	p.basicpay = basicPay
}

func (p *Permanent) ChangePf(pf int) {
	p.pf = pf
}

func (p *Permanent) Clone() Employee {
	return &Permanent{
		empId: p.empId,
		basicpay: p.basicpay,
		pf: p.pf,
	}
}