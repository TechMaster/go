package test

import (
	"fiber-postgres/model"
	"fiber-postgres/repo"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func Test_GetPostsOfUser(t *testing.T) {
	posts, err := repo.GetPostsOfUser("4Mw944HY")
	assert.Nil(t, err)

	// In ra thông tin danh sách post của user
	for i := 0; i < len(posts); i++ {
		posts[i].Print()
	}
}

func Test_GetPostDetail(t *testing.T) {
	post, err := repo.GetPostDetail("4Mw944HY", "35RnZDz7")
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_CreatePost(t *testing.T) {
	req := &model.CreatePost{
		Title:    "New post 1",
		Content:  gofakeit.LoremIpsumSentence(200),
	}

	userId := "4Mw944HY"

	post, err := repo.CreatePost(userId, req)
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_UpdatePost(t *testing.T) {
	req := &model.CreatePost{
		Title:    "Post 1 update",
		Content:  gofakeit.LoremIpsumSentence(200),
	}

	userId := "4Mw944HY"
	postId := "pRRFPqgZ"

	post, err := repo.UpdatePost(userId, postId, req)
	assert.Nil(t, err)

	// In ra thông tin của post
	post.Print()
}

func Test_DeletePost(t *testing.T) {
	userId := "4Mw944HY"
	postId := "pRRFPqgZ"

	err := repo.DeletePost(userId, postId)
	assert.Nil(t, err)
}
