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

/*
Bài tập về nhà
1. Sửa tên 1 thành viên trong members
2. Sửa tên club
3. Xóa 1 thành viên
4. Kiểm tra mỗi thành viên tham gia bao nhiêu câu lạc bộ
5. Kiểm tra mỗi câu lạc bộ được tham gia bởi bao nhiêu thành viên
6. Xóa member theo id (hoặc name)
*/ 

