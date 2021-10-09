package repo

import (
	"many_many/model"

	"github.com/go-pg/pg/v10"
)

//https://github.com/eleanorhealth/milo
func Add_member_to_club() (err error) {
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}
	//---- Tạo members
	john := model.Member{
		Id:   NewID(),
		Name: "John"}

	anna := model.Member{
		Id:   NewID(),
		Name: "Anna"}

	bob := model.Member{
		Id:   NewID(),
		Name: "Bob"}

	alice := model.Member{
		Id:   NewID(),
		Name: "Alice"}

	_, err = transaction.Model(&john, &anna, &bob, &alice).Insert()
	if !check_err(err, transaction) {
		return err
	}

	//--- Club
	math := model.Club{
		Id:   NewID(),
		Name: "Math",
	}

	sport := model.Club{
		Id:   NewID(),
		Name: "Sport",
	}

	music := model.Club{
		Id:   NewID(),
		Name: "Music",
	}
	_, err = transaction.Model(&math, &sport, &music).Insert()
	if !check_err(err, transaction) {
		return err
	}

	//---- Thêm các thành viên vào club
	err = Assign_members_to_club(transaction, math, []model.Member{john, anna})
	if !check_err(err, transaction) {
		return err
	}

	err = Assign_members_to_club(transaction, sport, []model.Member{bob, alice})
	if !check_err(err, transaction) {
		return err
	}

	err = Assign_members_to_club(transaction, music, []model.Member{john, bob, alice})
	if !check_err(err, transaction) {
		return err
	}
	return transaction.Commit()
}

/*
Thêm thành viên vào một club
*/
func Assign_members_to_club(transaction *pg.Tx, club model.Club, members []model.Member) (err error) {
	for _, member := range members {
		_, err = transaction.Model(&model.MemberClub{
			MemberId: member.Id,
			ClubId:   club.Id,
			Active:   random.Intn(2) == 1, //random true or false
		}).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

// Tìm club theo tên, đồng thời lấy danh sách thành viên
func Get_club_by_name(name string) (club *model.Club, err error) {
	club = new(model.Club)
	err = DB.Model(club).Relation("Members").Where("name ILIKE ?", name).Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return club, nil
}

// Tìm thành viên, lấy danh sách club mà người đó tham gia
func Get_member_by_name(name string) (member *model.Member, err error) {
	member = new(model.Member)
	err = DB.Model(member).Relation("Clubs").Where("name ILIKE ?", name).Limit(1).Select()
	if err != nil {
		return nil, err
	}
	return member, nil
}

type Club_Membership struct {
	Club   string `json:"club"`
	Status bool   `json:"status"`
}
type Member_Clubs struct {
	Id    string
	Name  string
	Clubs []Club_Membership
}

/*
Khi muốn lấy thêm trường bổ xung ở bảng trung gian. Ví dụ liệt kê tất cả các thành viên.
Trong mỗi thành viên liệt kê club họ tham gia và trạng thái tham gia
John
   Math : false
   Music : true
Anna
   Math : true
Bob
   Sport : true
   Music : false
*/
func Get_active_members_of_club() (members []Member_Clubs, err error) {
	query := `SELECT m2.id, m2.name, result.clubs FROM test."member" m2
	INNER JOIN
	(SELECT m.id,
	json_agg(
	json_build_object('club', c.name, 'status', mc.active)) clubs
	FROM test."member" m
	INNER JOIN test.member_club mc ON m.id = mc.member_id
	INNER JOIN test.club c ON mc.club_id = c.id
	GROUP BY m.id) result ON m2.id = result.id`

	_, err = DB.Query(&members, query)
	if err != nil {
		return nil, err
	} else {
		return members, nil
	}
}
