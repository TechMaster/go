package test

import (
	"fiber-postgres/model"
	"fiber-postgres/repo"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func Test_GetPostsOfUser(t *testing.T) {
	posts, err := repo.GetPostsOfUser("degz_G9v")
	assert.Nil(t, err)

	// In ra thông tin danh sách post của user
	for i := 0; i < len(posts); i++ {
		posts[i].Print()
	}
}

func Test_GetPostDetail(t *testing.T) {
	post, err := repo.GetPostDetail("degz_G9v", "W3k_aV4F")
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_CreatePost(t *testing.T) {
	req := &model.CreatePost{
		Title:    "New post 3",
		Content:  gofakeit.LoremIpsumSentence(200),
	}

	userId := "degz_G9v"

	post, err := repo.CreatePost(userId, req)
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_UpdatePost(t *testing.T) {
	req := &model.CreatePost{
		Title:    "Post 3 update",
		Content:  gofakeit.LoremIpsumSentence(200),
	}

	userId := "degz_G9v"
	postId := "7mE04tSE"

	post, err := repo.UpdatePost(userId, postId, req)
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_DeletePost(t *testing.T) {
	userId := "degz_G9v"
	postId := "7mE04tSE"

	err := repo.DeletePost(userId, postId)
	assert.Nil(t, err)
}
