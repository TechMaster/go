package repo

import (
	"user_role/model"
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
Chuyển đổi int[] thành string[]
*/
func convert_introles_to_enumroles(int_roles []int) (enum_roles []string) {
	for _, role := range int_roles {
		enum_roles = append(enum_roles, model.ROLES[role])
	}
	return enum_roles
}
