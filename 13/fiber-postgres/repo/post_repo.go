package repo

import (
	"errors"
	"fiber-postgres/model"
	"time"
)

/* Lấy danh sách post của user */
func GetPostsOfUser(userId string) (posts []model.Post, err error) {
	_, err = DB.Query(&posts, `
		SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at
		FROM test.posts p
		INNER JOIN test.users u
		ON p.author_id = u.id
		WHERE u.id = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	return posts, err
}

/* Lấy chi tiết bài post */
func GetPostDetail(userId string, postId string) (post model.Post, err error) {
	_, err = DB.Query(&post, `
		SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at
		FROM test.posts p
		INNER JOIN test.users u
		ON p.author_id = u.id
		WHERE p.author_id = ? AND p.id = ?
	`, userId, postId)

	if err != nil {
		return model.Post{}, err
	}

	return post, err
}

/* Tạo post mới */
func CreatePost(userId string, req *model.CreatePost) (post *model.Post, err error) {
	if req.Title == "" {
		return nil, errors.New("tiêu đề không được để trống")
	}

	post = &model.Post{
		Id:        NewID(),
		Title:     req.Title,
		Content:   req.Content,
		AuthorId:  userId,
		CreatedAt: time.Now(),
	}
	_, err = DB.Model(post).WherePK().Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return post, nil
}

/* Cập nhật thông tin post */
func UpdatePost(userId string, postId string, req *model.CreatePost) (post *model.Post, err error) {
	if req.Title == "" {
		return nil, errors.New("tiêu đề không được để trống")
	}

	post = &model.Post{
		Id:        postId,
		Title:     req.Title,
		Content:   req.Content,
		AuthorId:  userId,
		UpdatedAt: time.Now(),
	}
	_, err = DB.Model(post).Column("title", "content", "updated_at").Returning("*").Where("id = ? AND author_id = ?", postId, userId).Update()
	if err != nil {
		return nil, err
	}

	return post, nil
}

/* Xóa post */
func DeletePost(userId string, postId string) (err error) {
	_, err = DB.Exec(`
		DELETE FROM test.posts
		WHERE id = ? AND author_id = ?
	`, postId, userId)
	if err != nil {
		return err
	}

	return nil
}
