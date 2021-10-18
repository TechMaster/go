package test

import (
	"fiber-postgres/model"
	"fiber-postgres/repo"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllUser(t *testing.T) {
	users, err := repo.GetAllUser()
	assert.Nil(t, err)

	// In ra thông tin của 5 user
	for i := 0; i < 5; i++ {
		users[i].Print()
	}
}

func Test_GetUserById(t *testing.T) {
	user, err := repo.GetUserById("OwhzVBMO")
	assert.Nil(t, err)

	// In ra thông tin của user
	user.Print()
}

func Test_CreateUser(t *testing.T) {
	req := &model.CreateUser{
		FullName: "Bui Van Hien",
		Phone:    gofakeit.Phone(),
		Email:    gofakeit.Email(),
	}

	user, err := repo.CreateUser(req)
	assert.Nil(t, err)

	// In ra thông tin của user
	user.Print()
}

func Test_UpdateUser(t *testing.T) {
	req := &model.CreateUser{
		FullName: "Hien",
		Phone:    gofakeit.Phone(),
		Email:    gofakeit.Email(),
	}

	user, err := repo.UpdateUser("CPFiKLvh", req)
	assert.Nil(t, err)

	// In ra thông tin của user
	user.Print()
}

func Test_DeleteUser(t *testing.T) {
	err := repo.DeleteUser("CPFiKLvh")
	assert.Nil(t, err)
}
