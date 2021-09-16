package employee

type Sale struct {
	Employee Employee
}

func (s *Sale) SetName() BuildEmployee {
	s.Employee.Name = "Phạm Thị Mẫn"
	return s
}

func (s *Sale) SetEmail() BuildEmployee {
	s.Employee.Email = "man@techmaster.vn"
	return s
}

func (s *Sale) SetAge() BuildEmployee {
	s.Employee.Age = 25
	return s
}

func (s *Sale) SetPosition() BuildEmployee {
	s.Employee.Position = "Sale"
	return s
}

func (s *Sale) GetEmployee() Employee {
	return s.Employee
}