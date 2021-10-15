## Quan hệ One To Many (quan hệ một - nhiều)

Ví dụ thực tế:

-  Một công ty có nhiều nhân viên
-  Một lớp học có nhiều học viên
-  Một thư viện có nhiều đầu sách
-  ...

### Định nghĩa model

```go
type Bar struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	FooId string `gorm:"column:foo_id"`
	Foo   Foo
}

type Foo struct {
	Id   string `gorm:"primaryKey"`
	Name string
	Bars []Bar `gorm:"foreignKey:FooId"`
}
```


### Chỉ định tên bảng

Chúng ta có thể chỉ định tên bảng ứng với các struct đã định nghĩa nếu không muốn sử dụng tên bảng mặc định do GORM sinh ra

Link tham khảo : [https://gorm.io/docs/conventions.html](https://gorm.io/docs/conventions.html)

```go
func(b *Bar) TableName() string {
	return "bar"
}

func(b *Foo) TableName() string {
	return "foo"
}
```

### Thực hành

Học viên thực hành các ví dụ trong [repo/query](./repo/query.go)