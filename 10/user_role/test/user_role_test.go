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

func Test_find_user_by_name_or_email(t *testing.T) {
	name_or_email := "jermainerosenbaum@koss.org"
	user, err := repo.Find_user_by_name_or_email(name_or_email)
	assert.Nil(t, err)
	//assert.Equal(t, name_or_email, user.Name)
	assert.Equal(t, name_or_email, user.Email)
}

func Test_find_user_has_roles(t *testing.T) {
	roles := []int{5, 4, 8}
	user, err := repo.Find_user_has_roles(roles...)
	fmt.Println(user)
	assert.Nil(t, err)
}

func Test_find_user_has_any_role(t *testing.T) {
	roles := []int{5, 4, 8}
	user, err := repo.Find_user_has_any_role(roles...)
	fmt.Println(user)
	assert.Nil(t, err)
}

func Test_find_shared_roles_of_users(t *testing.T) {
	roles, err := repo.Find_shared_roles_of_users("9O-vtkRd", "UcSg4ZyQ")
	fmt.Println(roles)
	assert.Nil(t, err)
}

func Test_add_role_to_user(t *testing.T) {
	roles := []int{5, 4, 8}
	err := repo.Add_role_to_user("_1jtuhNr", roles)
	assert.Nil(t, err)
}

func Test_remove_role_from_user(t *testing.T) {
	err := repo.Remove_role_from_user("_1jtuhNrs", 6)
	assert.Nil(t, err)
}
