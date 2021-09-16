package employee

type Teacher struct {
	Employee Employee
}

func (t *Teacher) SetName() BuildEmployee {
	t.Employee.Name = "Nguyễn Xuân Ba"
	return t
}

func (t *Teacher) SetEmail() BuildEmployee {
	t.Employee.Email = "ba@techmaster.vn"
	return t
}

func (t *Teacher) SetAge() BuildEmployee {
	t.Employee.Age = 28
	return t
}

func (t *Teacher) SetPosition() BuildEmployee {
	t.Employee.Position = "Teacher"
	return t
}

func (t *Teacher) GetEmployee() Employee {
	return t.Employee
}
