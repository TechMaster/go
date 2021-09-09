package value

import (
	"github.com/brianvoe/gofakeit/v6"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Account struct {
	Id       string //Uniqe Id
	Email    string
	FullName string
	Address  gofakeit.AddressInfo
}

type AccountRepo struct {
	Accounts []Account //Che dấu dữ liệu tốt. Encapsulation in OOP
}

var AccRepo = AccountRepo{}

func InitData() {
	for i := 0; i < 100; i++ {
		id, _ := gonanoid.New(8)
		AccRepo.Accounts = append(AccRepo.Accounts, Account{
			Id:       id,
			Email:    gofakeit.Email(),
			FullName: gofakeit.Name(),
			Address:  *gofakeit.Address(),
		})
	}
}

func FindById(id string) *Account {
	for _, a := range AccRepo.Accounts {
		if a.Id == id {
			return &a
		}
	}
	return nil
}
