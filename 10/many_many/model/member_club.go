package model

//----- Quan hệ Nhiều - Nhiều
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
