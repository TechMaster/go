package repo

import (
	"errors"
	"fiber-postgres/model"
	"time"
)

/* Lấy danh sách user */
func GetAllUser() (users []model.User, err error) {
	err = DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}

/* Lấy thông tin user */
func GetUserById(id string) (user *model.User, err error) {
	user = &model.User{
		Id: id,
	}

	err = DB.Model(user).WherePK().Select()
	if err != nil {
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
	_, err = DB.Model(user).WherePK().Returning("*").Insert()
	if err != nil {
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
		Id:        id,
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		UpdatedAt: time.Now(),
	}

	_, err = DB.Model(user).Column("full_name", "phone", "email", "updated_at").Returning("*").WherePK().Update()
	if err != nil {
		return nil, err
	}

	return user, nil
}

/* Xóa user */
func DeleteUser(id string) (err error) {
	user := model.User{
		Id: id,
	}

	_, err = DB.Model(&user).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
