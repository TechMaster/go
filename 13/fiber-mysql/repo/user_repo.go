package repo

import (
	"errors"
	"fiber-postgres/model"
	"time"
)

/* Lấy danh sách user */
func GetAllUser() (users []model.User, err error) {
	if err = DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

/* Lấy thông tin user */
func GetUserById(id string) (user *model.User, err error) {
	if err = DB.Find(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

/* Tạo user */
func CreateUser(req *model.CreateUser) (user *model.User, err error) {
	if req.FullName == "" {
		return nil, errors.New("tên không được để trống")
	}

	user = &model.User{
		Id:        NewID(),
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
	}
	if err = DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

/* Cập nhật thông tin user */
func UpdateUser(id string, req *model.CreateUser) (user *model.User, err error) {
	if req.FullName == "" {
		return nil, errors.New("tên không được để trống")
	}

	user = &model.User{
		Id: id,
	}

	err = DB.Model(user).
		Select("full_name", "phone", "email", "updated_at").
		Updates(model.User{
			FullName:  req.FullName,
			Email:     req.Email,
			Phone:     req.Phone,
			UpdatedAt: time.Now(),
		}).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

/* Xóa user */
func DeleteUser(id string) (err error) {
	if err = DB.Delete(&model.User{}, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}
