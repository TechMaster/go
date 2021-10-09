package repo

import (
	"gopgdemo/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-pg/pg/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/*
Sinh ngẫu nhiên các roles trả về mảng int và string.
Mảng string dùng cho enum
*/
func gen_random_roles(numberOfRoles int) ([]int, []string) {
	int_roles := []int{}
	enum_roles := []string{}
	for i := 0; i < numberOfRoles; i++ {
		var role int
		for { //Loop cho đến khi tạo ra phần tử mới
			role = 1 + random.Intn(8)
			if !int_in_array(role, int_roles) {
				break
			}
		}

		int_roles = append(int_roles, role)
		enum_roles = append(enum_roles, model.ROLES[role])
	}
	return int_roles, enum_roles
}

/*
Tạo fake user với random name, email, di động, roles
*/
func Create_fake_user() error {
	var err error
	id, _ := gonanoid.New(8)
	int_roles, enum_roles := gen_random_roles(3)

	transaction, err := DB.Begin()
	if err != nil {
		return err
	}

	user := model.User{
		Id:         id,
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Mobile:     gofakeit.Phone(),
		Int_roles:  int_roles,
		Enum_roles: enum_roles,
	}

	_, err = transaction.Model(&user).Insert()
	if !check_err(err, transaction) {
		return err
	}

	for _, role := range int_roles {
		user_role := model.User_Role{
			User_id: id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !check_err(err, transaction) {
			return err
		}
	}

	err = transaction.Commit()
	return err
}

/*
Bổ xung user với tên, email, mobile, role thật
*/
func create_user(transaction *pg.Tx, name string, email string, mobile string, roles []int) (err error) {
	user_id := NewID()
	user := model.User{
		Id:        user_id,
		Name:      name,
		Email:     email,
		Mobile:    mobile,
		Int_roles: roles,
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

/*
Tìm kiếm user theo tên hoặc email
Linh làm
*/
func GetUserByNameOrEmail(name_or_email string) *model.User {
	return nil
}
