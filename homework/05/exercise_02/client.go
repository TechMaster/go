package exercise_02

import "fmt"

func DemoExercise_02() {
    // Tạo đối tượng Permanent
    p := Permanent{
        empId: 1,
        basicpay: 1000,
        pf: 100,
    }

    p1 := p.Clone()
    p1.ChangeEmId(2)
    p1.ChangeBasicPay(2000)
    p1.ChangePf(200)

    p2 := p.Clone()
    p2.ChangeEmId(3)
    p2.ChangeBasicPay(3000)
    p2.ChangePf(300)

    // Tạo đối tượng Contact
    c := Contract{
        empId: 4,
        basicpay: 500,
    }

    c1 := c.Clone()
    c1.ChangeEmId(5)
    c1.ChangeBasicPay(600)

    // Tính tổng tiền
    employees := []Employee{&p, p1, p2, &c, c1}
    fmt.Println(CalculateSalary(employees))
}

// Hàm tính tổng tiền cho một danh sách employee
func CalculateSalary(employees []Employee) int {
    total := 0
    for i := 0; i < len(employees); i++ {
        total += employees[i].GetSalary()
    }
    return total
}




