package test

import (
	"fmt"
	"testing"
	"user_role/repo"

	"github.com/stretchr/testify/assert"
)

func Test_create_fake_user(t *testing.T) {
	user, err := repo.Create_fake_user()
	assert.Nil(t, err)
	if err == nil {
		fmt.Println(user.Id, user.Name)
	}
}

func Test_create_400_fake_user(t *testing.T) {
	transaction, err := repo.DB.Begin()
	assert.Nil(t, err)
	if err != nil {
		return
	}

	for i := 0; i < 400; i++ {
		user, err := repo.Create_fake_user(transaction)
		assert.Nil(t, err)

		if err != nil {
			_ = transaction.Rollback()
			return
		} else {
			fmt.Printf("%d - %s - %s\n", i, user.Id, user.Name)
		}
	}
	err = transaction.Commit()
	assert.Nil(t, err)
}
