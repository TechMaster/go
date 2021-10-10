package repo

import (
	"user_role/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-pg/pg/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/*
Tạo fake user với random name, email, di động, roles
Nếu có tham số trans thì dùng tham số trans làm transaction, còn không tự tạo mới transaction
*/
func Create_fake_user(trans ...*pg.Tx) (user *model.User, err error) {
	var transaction *pg.Tx
	if len(trans) > 0 {
		transaction = trans[0]
	} else {
		transaction, err = DB.Begin()
		if err != nil {
			return nil, err
		}
	}

	id, _ := gonanoid.New(8)
	int_roles, enum_roles := gen_random_roles(3)

	user = &model.User{
		Id:         id,
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Mobile:     gofakeit.Phone(),
		Int_roles:  int_roles,
		Enum_roles: enum_roles,
	}

	_, err = transaction.Model(user).Insert()
	if !check_err(err, transaction) {
		return nil, err
	}

	for _, role := range int_roles {
		user_role := model.User_Role{
			User_id: id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !check_err(err, transaction) {
			return nil, err
		}
	}

	//Nếu dùng transaction tạo trong hàm này thì commit luôn
	if len(trans) == 0 {
		if err = transaction.Commit(); err != nil {
			return nil, err
		} else {
			return user, nil
		}
	} else {
		return user, nil
	}
}

/*
Bổ xung user với tên, email, mobile, role thật
*/
func Create_user(transaction *pg.Tx, name string, email string, mobile string, roles []int) (err error) {
	user_id := NewID()
	user := model.User{
		Id:         user_id,
		Name:       name,
		Email:      email,
		Mobile:     mobile,
		Int_roles:  roles,
		Enum_roles: convert_introles_to_enumroles(roles),
	}

	_, err = transaction.Model(&user).Insert()
	if !check_err(err, transaction) {
		return err
	}

	for _, role := range roles {
		user_role := model.User_Role{
			User_id: user_id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !check_err(err, transaction) {
			return err
		}
	}
	return nil
}
