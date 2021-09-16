package post_factory

type post struct {
	Title string
	Author string
}

/* function Get Title */
func (p *post) GetTitle() string {
	return p.Title
}

func (p *post) GetAuthor() string {
	return p.Author
}

func (p *post) SetAuthor(author string) {
	p.Author = author
}

func (p *post) SetTitle(author string) {
	p.Title = author
}
