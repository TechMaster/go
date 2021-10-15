## Quan hệ Many To Many (quan hệ nhiều - nhiều)

Ví dụ thực tế:

-  Một lớp có thể có nhiều sinh viên, một sinh viên có thể học nhiều lớp
-  Một người dùng có thể mua nhiều mặt hàng, một mặt hàng có thể được mua bởi nhiều người

### Định nghĩa model

```go
type Member struct {
	Id    string `gorm:"primaryKey"`
	Name  string
	Clubs []Club `gorm:"many2many:member_club;"`
}

type Club struct {
	Id      string `gorm:"primaryKey"`
	Name    string
	Members []Member `gorm:"many2many:member_club;"`
}

type MemberClub struct {
	MemberId string `gorm:"primaryKey" column:"member_id"`
	ClubId   string `gorm:"primaryKey" column:"club_id"`
	Active   bool
}
```


### Chỉ định tên bảng

Chúng ta có thể chỉ định tên bảng ứng với các struct đã định nghĩa nếu không muốn sử dụng tên bảng mặc định do GORM sinh ra

```go
func(m *Member) TableName() string {
	return "member"
}

func(c *Club) TableName() string {
	return "club"
}

func(mc *MemberClub) TableName() string {
	return "member_club"
}
```

### Thực hành

Học viên thực hành các ví dụ trong [repo/query](./repo/query.go)
