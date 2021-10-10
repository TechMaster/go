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
		fmt.Printf("%s - %d - %s\n", web_track.Id, web_track.Version, web_track.Name)
	}
}
