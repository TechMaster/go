package model

type Enum_roles_int int

const (
	ROLE_ADMIN      Enum_roles_int = 1
	ROLE_STUDENT                   = 2
	ROLE_TRAINER                   = 3
	ROLE_SALE                      = 4
	ROLE_EMPLOYER                  = 5
	ROLE_AUTHOR                    = 6
	ROLE_EDITOR                    = 7 //edit bài, soạn page, làm công việc digital marketing
	ROLE_MAINTAINER                = 8 //quản trị hệ thống, gánh bớt việc cho Admin, back up dữ liệu. Sửa đổi profile,role user, ngoại trừ role ROOT và Admin
)

var (
	ROLES = []string{"", "ADMIN", "STUDENT", "TRAINER", "SALE", "EMPLOYER", "AUTHOR", "EDITOR", "MAINTAINER"}
)
