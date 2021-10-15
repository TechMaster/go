package model

import (
	"basic/util"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

/*
HOOK
Hook là các hàm được gọi trước hoặc sau khi thực hiện các thao tác creation, querying, updating, deletion

Chúng ta có thể chỉ định hook cho một model (struct) cụ thể nào đó, nó sẽ tự động được gọi khi creation, querying, updating, deletion

Nếu khi hook được gọi mà xảy ra lỗi, lúc này GORM sẽ rollback transaction
*/

// Create HOOK
// BeforeCreate, BeforeSave, AfterCreate, AfterSave

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate")
	if s.Id == "" {
		fmt.Println("Bổ sung Id cho student")
		s.Id = util.NewID()
	}

	return
}

// Delete
func (s *Student) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete")
	if s.Id == "1" {
		return errors.New("user not allowed to delete")
	}
	return
}

// Query hook
func (s *Student) AfterFind(tx *gorm.DB) (err error) {
	s.ToString()
	return
}
