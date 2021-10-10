package repo

import "github.com/go-pg/pg/v10"

/*
Thêm nhiều role vào cho user
Role nào user đã có rồi thì bỏ qua, chỉ thêm những role user chưa có
*/
func Add_role_to_user(user_id string, roles []int) (err error) {
	_, err = DB.Query(1, `update test.users set int_roles = (select array_agg(distinct e) from unnest(int_roles || ?0) e) where id = ?1
	and not int_roles @> ?0`, pg.Array(roles), user_id)
	return err
}

/*
Loại bỏ nhiều role ra khỏi một user
Role nào user không có thì bỏ qua, chỉ xoá những role người dùng có
*/
func Remove_role_from_user(user_id string, role int) (err error) {
	_, err = DB.Query(1, `update test.users set int_roles = array_remove(int_roles, ?0) where id = ?1`,
		role, user_id)
	return err
}
