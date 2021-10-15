package test

import (
	"fmt"
	"many-to-many/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddMemberToClub(t *testing.T) {
	db := repo.GetInstance()
	err := repo.AddMemberToClub(db)
	assert.Nil(t, err)
}

func Test_GetClubByName(t *testing.T) {
	db := repo.GetInstance()
	club, err := repo.GetClubByName(db, "Sport")
	assert.Nil(t, err)

	fmt.Println(club)

	fmt.Println(club.Name)
	for _, m := range club.Members {
		fmt.Println("	" + m.Name)
	}
}

func Test_GetMemberByName(t *testing.T) {
	db := repo.GetInstance()
	member, err := repo.GetMemberByName(db, "Bob")
	assert.Nil(t, err)

	fmt.Println(member)

	fmt.Println(member.Name)
	for _, c := range member.Clubs {
		fmt.Println("	" + c.Name)
	}
}
