# Quan hệ Nhiều-Nhiều
Bảng trung gian có thêm cột `active`

## Định nghĩa model
```go
type Member struct {
	tableName struct{} `pg:"test.member"`
	Id        string   `pg:"id,pk"`
	Name      string
	Clubs     []Club `pg:"many2many:test.member_club"`
}

type Club struct {
	tableName struct{} `pg:"test.club"`
	Id        string   `pg:"id,pk"`
	Name      string
	Members   []Member `pg:"many2many:test.member_club"` //Phải bỏ rel mà chỉ dùng many2many thôi
}

type MemberClub struct {
	tableName struct{} `pg:"test.member_club"`
	MemberId  string   `pg:"member_id,pk"`
	ClubId    string   `pg:"club_id,pk"`
	Active    bool     `pg:",use_zero"` //Bắt buộc phải dùng pg:",use_zero" nếu không giá trị false biến thành null
}

```

## Đăng ký ORM

```go
orm.RegisterTable(&model.MemberClub{})
```

## Bài tập

1. Xoá một thành viên ra khỏi club
2. Thay đổi trạng thái membership