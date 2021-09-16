package student

type studentBuilder struct {
    Name string
    Sex   string
    Age      int
}

func newStudentBuilder() *studentBuilder {
    return &studentBuilder{}
}

func (b *studentBuilder) setName(name string) iBuilder {
    b.Name = name
    return b
}

func (b *studentBuilder) setSex(sex string) iBuilder{
    b.Sex = sex
    return b
}

func (b *studentBuilder) setAge(age int) iBuilder{
    b.Age = age
    return b
}

func (b *studentBuilder) getStudent() student {
    return student{
        Name: b.Name,
        Sex: b.Sex,
        Age: b.Age,
    }
}
