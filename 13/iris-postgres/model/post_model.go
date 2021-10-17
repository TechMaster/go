package model

import (
	"fmt"
	"time"
)

type Post struct {
	tableName struct{}  `pg:"test.posts"`
	Id        string    `pg:"id,pk" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorId  string    `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// User      User `pg:"rel:has-one"`
}

/*
Struct này dùng để create và update post
*/
type CreatePost struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
}

/*
In ra thông tin của post
*/
func (p *Post) Print() {
	fmt.Println("Id : ", p.Id)
	fmt.Println("Title : ", p.Title)
	fmt.Println("Content : ", p.Content)
	fmt.Println("AuthorId : ", p.AuthorId)
	fmt.Println("CreatedAt : ", p.CreatedAt)
	fmt.Println("UpdatedAt : ", p.UpdatedAt)
	fmt.Println()
}
