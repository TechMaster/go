package post_factory

type newsPost struct {
	post
}

func newNewsPost() IPost {
	return &newsPost {
		post : post{
			Title: "Tin tức buổi sáng",
			Author : "Nguyễn Thu Hằng",
		},
	}
}