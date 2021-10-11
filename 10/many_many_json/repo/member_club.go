package repo

import (
	"many_many_json/model"
	"reflect"

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
	err = Assign_members_to_club(transaction, &math, []*model.Member{&john, &anna})
	if !check_err(err, transaction) {
		return err
	}

	err = Assign_members_to_club(transaction, &sport, []*model.Member{&bob, &alice})
	if !check_err(err, transaction) {
		return err
	}

	err = Assign_members_to_club(transaction, &music, []*model.Member{&john, &bob, &alice})
	if !check_err(err, transaction) {
		return err
	}
	return transaction.Commit()
}

/*
Thêm thành viên vào một club
*/
func Assign_members_to_club(transaction *pg.Tx, club *model.Club, members []*model.Member) (err error) {
	for _, member := range members {
		membership := model.MemberShip{
			MemberId: member.Id,
			ClubId:   club.Id,
			IsActive: random.Intn(2) == 1, //random true or false
		}
		upsert_membership(&club.Memberships, membership)
		_, err = transaction.Model(club).Column("memberships").WherePK().Update()
		if !check_err(err, transaction) {
			return err
		}
		upsert_membership(&member.Memberships, membership)
		_, err = transaction.Model(member).Column("memberships").WherePK().Update()
		if !check_err(err, transaction) {
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

func upsert_membership(memberships *[]model.MemberShip, new_member_ship model.MemberShip) {
	isNewMembership := true
	if *memberships == nil {
		*memberships = []model.MemberShip{new_member_ship}
		return
	}

	for _, membership := range *memberships {
		if reflect.DeepEqual(membership, new_member_ship) {
			isNewMembership = false
			break
		}

		//Trùng clubid và memberid nhưng khác trạng thái IsActive
		if membership.ClubId == new_member_ship.ClubId && membership.MemberId == new_member_ship.MemberId {
			membership.IsActive = new_member_ship.IsActive
			isNewMembership = false
			break
		}
	}

	if isNewMembership {
		*memberships = append(*memberships, new_member_ship)
	}

}
