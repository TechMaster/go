package employee

type BuildEmployee interface {
	SetName() BuildEmployee
	SetEmail() BuildEmployee
	SetAge() BuildEmployee
	SetPosition() BuildEmployee
	GetEmployee() Employee
}
