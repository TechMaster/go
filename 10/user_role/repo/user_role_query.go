package repo

import (
	"fmt"
	"user_role/model"
)

/*
Tìm kiếm user theo tên hoặc email
*/
func Find_user_by_name_or_email(name_or_email string) (user model.User, err error) {
	//Dùng raw query
	_, err = DB.Query(&user, `SELECT * FROM test.users u WHERE u.name= ?0 or u.email= ?0`, name_or_email)
	return user, err
}

/*
Tìm tất cả các user tất cả các role. Sử dụng AND
*/
func Find_user_has_roles(roles ...int) (users []model.User, err error) {
	query := "SELECT * FROM test.users u WHERE "

	for i := 0; i < len(roles); i++ {
		if i != len(roles)-1 {
			query += fmt.Sprintf(" ?%d = ANY(u.int_roles) AND", i)
		} else {
			query += fmt.Sprintf(" ?%d = ANY(u.int_roles)", i)
		}
	}

	params := make([]interface{}, len(roles))
	for i, role := range roles {
		params[i] = role
	}

	_, err = DB.Query(&users, query, params...)

	return users, err
}

/*
Tìm tất cả các user có một trong các role sau đây. Sử dụng OR
*/
func Find_user_has_any_role(roles ...int) (users []model.User, err error) {
	query := "SELECT * FROM test.users u WHERE "

	for i := 0; i < len(roles); i++ {
		if i != len(roles)-1 {
			query += fmt.Sprintf(" ?%d = ANY(u.int_roles) OR", i)
		} else {
			query += fmt.Sprintf(" ?%d = ANY(u.int_roles)", i)
		}
	}

	params := make([]interface{}, len(roles))
	for i, role := range roles {
		params[i] = role
	}

	_, err = DB.Query(&users, query, params...)

	return users, err
}

/*
Xác định 2 user có chung những role nào
*/
func Find_shared_roles_of_users(userA_id, userB_id string) (roles []int, err error) {
	_, err = DB.Query(&roles, `SELECT unnest(u.int_roles) FROM test.users u WHERE u.id = ?0 or u.id = ?1 
	group by 1
	having count(*) > 1;`, userA_id, userB_id)
	return roles, err
}
