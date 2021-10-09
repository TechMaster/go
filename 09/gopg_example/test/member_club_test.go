package test

import (
	"fmt"
	"gopgdemo/repo"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_member_club(t *testing.T) {
	err := repo.AddMemberClub()
	assert.Nil(t, err)
}

func Test_get_club_by_name(t *testing.T) {
	club, err := repo.Get_club_by_name("math")
	assert.Nil(t, err)
	if err == nil {
		for _, member := range club.Members {
			fmt.Println(member.Name)
		}
	}
}

func Test_get_member_by_name(t *testing.T) {
	member, err := repo.Get_member_by_name("bob")
	assert.Nil(t, err)
	if err == nil {
		for _, club := range member.Clubs {
			fmt.Println(club.Name)
		}
	}
}

func Test_get_active_members_of_club(t *testing.T) {
	members, err := repo.Get_active_members_of_club()
	assert.Nil(t, err)
	if err == nil {
		for _, member := range members {
			fmt.Println(member.Name)
			for _, membership := range member.Clubs {
				fmt.Println("   " + membership.Club + " : " + strconv.FormatBool(membership.Status))
			}
		}
	}
}
