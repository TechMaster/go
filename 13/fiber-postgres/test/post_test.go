package test

import (
	"fiber-postgres/model"
	"fiber-postgres/repo"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func Test_GetPostsOfUser(t *testing.T) {
	posts, err := repo.GetPostsOfUser("j2B5GYRa")
	assert.Nil(t, err)

	// In ra thông tin danh sách post của user
	for i := 0; i < len(posts); i++ {
		posts[i].Print()
	}
}

func Test_GetPostDetail(t *testing.T) {
	post, err := repo.GetPostDetail("j2B5GYRa", "cs0hSegU")
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_CreatePost(t *testing.T) {
	req := &model.CreatePost{
		Title:    "New post 3",
		Content:  gofakeit.LoremIpsumSentence(200),
	}

	userId := "j2B5GYRa"

	post, err := repo.CreatePost(userId, req)
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_UpdatePost(t *testing.T) {
	req := &model.CreatePost{
		Title:    "Post 2 update",
		Content:  gofakeit.LoremIpsumSentence(200),
	}

	userId := "j2B5GYRa"
	postId := "CUHDEesq"

	post, err := repo.UpdatePost(userId, postId, req)
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_DeletePost(t *testing.T) {
	userId := "j2B5GYRa"
	postId := "CUHDEesq"

	err := repo.DeletePost(userId, postId)
	assert.Nil(t, err)
}
