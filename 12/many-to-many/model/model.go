package model

//----- Quan hệ Nhiều - Nhiều
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

func(m *Member) TableName() string {
	return "member"
}

func(c *Club) TableName() string {
	return "club"
}

func(mc *MemberClub) TableName() string {
	return "member_club"
}
