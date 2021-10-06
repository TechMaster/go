package model

const (
	ADMIN      = 1
	STUDENT    = 2
	TRAINER    = 3
	SALE       = 4
	EMPLOYER   = 5
	AUTHOR     = 6
	EDITOR     = 7 //edit bài, soạn page, làm công việc digital marketing
	MAINTAINER = 8 //quản trị hệ thống, gánh bớt việc cho Admin, back up dữ liệu. Sửa đổi profile,role user, ngoại trừ role ROOT và Admin

	//---Status
	STATUS_DRAFT  = "draft"
	STATUS_ACTIVE = "active"
	STATUS_HIDDEN = "hidden"
)

var (
	ROLES = []string{"", "ADMIN", "STUDENT", "TRAINER", "SALE", "EMPLOYER", "AUTHOR", "EDITOR", "MAINTAINER"}
)
