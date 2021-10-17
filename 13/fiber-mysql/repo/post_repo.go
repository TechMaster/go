package repo

import (
	"errors"
	"fiber-postgres/model"
	"time"
)

/* Lấy danh sách post của user */
func GetPostsOfUser(userId string) (posts []model.Post, err error) {
	err = DB.Raw(`
		SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at
		FROM db.posts p
		INNER JOIN db.users u
		ON p.author_id = u.id
		WHERE u.id = ?
	`, userId).Scan(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, err
}

/* Lấy chi tiết bài post */
func GetPostDetail(userId string, postId string) (post model.Post, err error) {
	err = DB.Raw(`
		SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at
		FROM db.posts p
		INNER JOIN db.users u
		ON p.author_id = u.id
		WHERE p.author_id = ? AND p.id = ?
	`, userId, postId).Scan(&post).Error

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
	if err = DB.Create(post).Error; err != nil {
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
		Id:       postId,
		AuthorId: userId,
	}
	err = DB.Model(post).
		Select("title", "content", "updated_at").
		Updates(model.Post{
			Title:     req.Title,
			Content:   req.Content,
			UpdatedAt: time.Now(),
		}).
		Where("id = ? AND author_id = ?", postId, userId).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

/* Xóa post */
func DeletePost(userId string, postId string) (err error) {
	err = DB.Exec(`
		DELETE FROM db.posts
		WHERE id = ? AND author_id = ?
	`, postId, userId).Error
	if err != nil {
		return err
	}

	return nil
}
