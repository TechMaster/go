package post_factory

import "fmt"

func PostFactory() {
	news, _ := GetPost("news")
	normal, _ := GetPost("normal")

	printDetailsPost(news)
	printDetailsPost(normal)

	news.SetAuthor("Nguyễn Bích Ngọc")
	news.SetTitle("Tin nóng tuần qua")
	printDetailsPost(news)

	_, err := GetPost("abc")
    if err != nil {
        fmt.Println(err)
    }
}

func printDetailsPost(p IPost) {
    fmt.Printf("Title: %s \n", p.GetTitle())
    fmt.Printf("Author: %s \n\n", p.GetAuthor())
}