package exercise_02

type Employee interface {
	Clone() Employee
	ChangeEmId(int)
	ChangeBasicPay(int)
	ChangePf(int)
	GetSalary() int
}