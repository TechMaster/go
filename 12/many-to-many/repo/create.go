package repo

import (
	"many-to-many/model"
	"many-to-many/util"

	"gorm.io/gorm"
)

func AddMemberToClub(db *gorm.DB) (err error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	//---- Tạo members
	john := model.Member{
		Id:   util.NewID(),
		Name: "John",
	}

	anna := model.Member{
		Id:   util.NewID(),
		Name: "Anna",
	}

	bob := model.Member{
		Id:   util.NewID(),
		Name: "Bob",
	}

	alice := model.Member{
		Id:   util.NewID(),
		Name: "Alice",
	}

	var members []model.Member
	members = append(members, john, anna, bob, alice)

	if err := tx.Create(&members).Error; err != nil {
		tx.Rollback()
		return err
	}

	//--- Club
	math := model.Club{
		Id:   util.NewID(),
		Name: "Math",
	}

	sport := model.Club{
		Id:   util.NewID(),
		Name: "Sport",
	}

	music := model.Club{
		Id:   util.NewID(),
		Name: "Music",
	}
	var clubs []model.Club
	clubs = append(clubs, math, sport, music)

	if err := tx.Create(&clubs).Error; err != nil {
		tx.Rollback()
		return err
	}

	//---- Thêm các thành viên vào club
	err = assignMembersToClub(tx, math, []model.Member{john, anna})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = assignMembersToClub(tx, sport, []model.Member{bob, alice})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = assignMembersToClub(tx, music, []model.Member{john, bob, alice})
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func assignMembersToClub(tx *gorm.DB, club model.Club, members []model.Member) (err error) {
	for _, member := range members {
		err := tx.Create(&model.MemberClub{
			MemberId: member.Id,
			ClubId:   club.Id,
			Active:   random.Intn(2) == 1, //random true or false
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
