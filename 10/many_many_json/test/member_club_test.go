package test

import (
	"fmt"
	"many_many_json/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_member_club(t *testing.T) {
	err := repo.Add_member_to_club()
	assert.Nil(t, err)
}

func Test_get_club_by_name(t *testing.T) {
	club, err := repo.Get_club_by_name("math")
	assert.Nil(t, err)
	if err == nil {
		for _, membership := range club.Memberships {
			fmt.Printf("%s - %s - %t\n", membership.ClubId, membership.MemberId, membership.IsActive)
		}
	}
}

func Test_get_member_by_name(t *testing.T) {
	member, err := repo.Get_member_by_name("bob")
	assert.Nil(t, err)
	if err == nil {
		for _, membership := range member.Memberships {
			fmt.Printf("%s - %s - %t\n", membership.ClubId, membership.MemberId, membership.IsActive)
		}
	}
}
