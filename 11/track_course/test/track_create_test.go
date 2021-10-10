package test

import (
	"fmt"
	"testing"
	"track_course/repo"

	"github.com/stretchr/testify/assert"
)

/*
Cần chạy file truncate_track.sql để xoá tất cả bản ghi track và track_master
*/
func Test_create_new_track_master(t *testing.T) {
	web_track, err := repo.Create_new_track_master("Web Front End React 7 tháng", "Mô tả Web Front End React 7 tháng")
	assert.Nil(t, err)
	if err == nil {
		fmt.Printf("%s - %s - %s\n", web_track.Id, web_track.Version, web_track.Name)
	}

	java_track, err := repo.Create_new_track_master("Java Spring Boot Full Stack 8 tháng", "Mô tả Java Spring Boot Full Stack 8 tháng")
	assert.Nil(t, err)
	if err == nil {
		fmt.Printf("%s - %s - %s\n", java_track.Id, java_track.Version, java_track.Name)
	}

	ios_track, err := repo.Create_new_track_master("Lập trình di động IOS 4 tháng", "Mô tả lập trình di động IOS 4 tháng")
	assert.Nil(t, err)
	if err == nil {
		fmt.Printf("%s - %s - %s\n", ios_track.Id, ios_track.Version, ios_track.Name)
	}
}