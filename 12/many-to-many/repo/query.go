package repo

import (
	"many-to-many/model"

	"gorm.io/gorm"
)

// Tìm club theo tên, đồng thời lấy danh sách thành viên
func GetClubByName(db *gorm.DB, name string) (club model.Club, err error) {
	err = db.Preload("Members").Find(&club, "name = ?", name).Error
	if err != nil {
		return model.Club{}, err
	}
	return club, nil
}

// Tìm thành viên, lấy danh sách club mà người đó tham gia
func GetMemberByName(db *gorm.DB, name string) (member model.Member, err error) {
	err = db.Preload("Clubs").Find(&member, "name = ?", name).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}