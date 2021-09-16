package student

type iBuilder interface {
    setName(string) iBuilder
    setSex(string) iBuilder
    setAge(int) iBuilder
    getStudent() student
}

func GetBuilder() iBuilder {
    return &studentBuilder{}
}