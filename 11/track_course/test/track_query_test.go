package test

import (
	"fmt"
	"testing"
	"track_course/repo"

	"github.com/stretchr/testify/assert"
)

func Test_find_track_master_by_name(t *testing.T) {
	web_track, err := repo.Find_track_by_name("web front end")
	assert.Nil(t, err)
	if err == nil {
		fmt.Printf("%s - %s - %s\n", web_track.Id, web_track.Version, web_track.Name)
	}
}

func Test_find_track_by_id(t *testing.T) {
	ios_track_master, err := repo.Find_track_by_name("di động IOS")
	track_000_id := ios_track_master.Id + "000"

	ios_track_000, err := repo.Find_track_by_id(track_000_id)

	assert.Nil(t, err)
	if err == nil {
		fmt.Printf("%s - %s - %s\n", ios_track_000.Id, ios_track_000.Version, ios_track_000.Name)
	}
}
