package post_factory

type normalPost struct {
	post
}

func newNormalPost() IPost {
	return &normalPost {
		post : post{
			Title: "Lập trình và cuộc sống",
			Author : "Nguyễn Minh Duy",
		},
	}
}