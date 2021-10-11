package model

//----- Quan hệ Nhiều - Nhiều
type Member struct {
	tableName   struct{} `pg:"test.memberj"`
	Id          string   `pg:"id,pk"`
	Name        string
	Memberships []MemberShip
}

type Club struct {
	tableName   struct{} `pg:"test.clubj"`
	Id          string   `pg:"id,pk"`
	Name        string
	Memberships []MemberShip
}

type MemberShip struct {
	MemberId string
	ClubId   string
	IsActive bool
}
