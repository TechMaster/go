package repo

import (
	"user_role/model"
)

/*
Tìm kiếm user theo tên hoặc email
*/
func Find_user_by_name_or_email(name_or_email string) (user *model.User, err error) {
	return nil, nil
}

/*
Tìm tất cả các user tất cả các role. Sử dụng AND
*/
func Find_user_has_roles(roles ...int) (users []model.User, err error) {
	return nil, nil
}

/*
Tìm tất cả các user có một trong các role sau đây. Sử dụng OR
*/
func Find_user_has_any_role(roles ...int) (users []model.User, err error) {
	return nil, nil
}

/*
Xác định 2 user có chung những role nào
*/
func Find_shared_roles_of_users(userA_id, userB_id string) (roles []int, err error) {
	return nil, nil
}
