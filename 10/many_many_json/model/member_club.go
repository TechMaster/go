package model

//----- Quan hệ Nhiều - Nhiều
type Member struct {
	tableName struct{} `pg:"test.memberj"`
	Id        string   `pg:"id,pk"`
	Name      string
	Clubs     []ClubInfo
}

type Club struct {
	tableName struct{} `pg:"test.clubj"`
	Id        string   `pg:"id,pk"`
	Name      string
	Members   []MemberInfo
}

type ClubInfo struct {
	Id       string //Id of club
	Name     string //Name of clubb
	IsActive bool
}

type MemberInfo struct {
	Id       string //Id of member
	Name     string //Name of member
	IsActive bool
}
