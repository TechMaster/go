package class

import "fmt"

type Class struct {
    Name string
	Session int
	Teacher string
}

func (c *Class) Print() {
    fmt.Printf("Lớp: %s\nSố buổi học: %d\nGiảng viên: %s\n\n", c.Name, c.Session, c.Teacher)
}

func (c *Class) SetName(name string) {
    c.Name = name
}

func (c *Class) SetSession(session int) {
    c.Session = session
}

func (c *Class) SetTeacher(teacher string) {
    c.Teacher = teacher
}

func (c *Class) Clone() iClass {
    return &Class{
        Name: c.Name,
        Session: c.Session,
		Teacher: c.Teacher,
    }
}