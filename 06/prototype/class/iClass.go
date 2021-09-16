package class

type iClass interface {
    Print()
	SetName(string)
    SetSession(int)
    SetTeacher(string)
    Clone() iClass
}