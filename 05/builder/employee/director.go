package employee

type Director struct {
	builder BuildEmployee
}

func (d *Director) SetBuilder(b BuildEmployee) {
	d.builder = b
}

func (d *Director) BuildEmployee() Employee {
	d.builder.SetName().SetEmail().SetAge().SetPosition()
	return d.builder.GetEmployee()
}
