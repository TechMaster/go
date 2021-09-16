package worker

import "fmt"

type Worker struct{
	Name string
	Email string
	CardId int
	SalaryRatio int
}

// Một số hàm khởi tạo lên đối tượng

func NewWorker(name string, email string) *Worker{
	return &Worker{
		Name: name,
		Email: email,
	}
}

func NewWorker02(name string, email string, cardId int) *Worker{
	return &Worker{
		Name: name,
		Email: email,
		CardId: cardId,
	}
}

func NewWorker03(name string, email string, salaryRatio int, cardId int) *Worker{
	return &Worker{
		Name: name,
		Email: email,
		CardId: cardId,
		SalaryRatio: salaryRatio,
	}
}

func (w Worker) GetSalary() int {
	return w.SalaryRatio * 1000000
}

func DemoWorker() {
	_ = NewWorker("Nguyễn Xuân Ba", "ba@techmaster.vn")
	_ = NewWorker02("Nguyễn Xuân Ba", "ba@techmaster.vn", 12345678)
	worker := NewWorker03("Nguyễn Xuân Ba", "ba@techmaster.vn", 12345678, 13)

	fmt.Println(worker.GetSalary())
}





