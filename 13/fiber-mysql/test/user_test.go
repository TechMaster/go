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
	user, err := repo.GetUserById("4Mw944HY")
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

	user, err := repo.UpdateUser("9--_58z6", req)
	assert.Nil(t, err)

	// In ra thông tin của user
	user.Print()
}

func Test_DeleteUser(t *testing.T) {
	err := repo.DeleteUser("9--_58z6")
	assert.Nil(t, err)
}

func Benchmark_GetAllUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = repo.GetAllUser() // 1104, 1093, 1077
	}
}

func Benchmark_GetUserById(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = repo.GetUserById("6jflpdcT") // 541, 574, 589
	}
}

func Benchmark_CreateUser(b *testing.B) {
	req := &model.CreateUser{
		FullName: gofakeit.Animal(),
		Phone: gofakeit.Phone(),
		Email: gofakeit.Email(),
	}

	for i := 0; i < b.N; i++ {
		_, _ = repo.CreateUser(req) // 206, 219, 217
	}
}

func Benchmark_UpdateUser(b *testing.B) {
	req := &model.CreateUser{
		FullName: gofakeit.Animal(),
		Phone: gofakeit.Phone(),
		Email: gofakeit.Email(),
	}

	for i := 0; i < b.N; i++ {
		_, _ = repo.UpdateUser("6jflpdcT", req) // 308, 288, 306
	}
}
