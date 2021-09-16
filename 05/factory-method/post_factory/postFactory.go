package post_factory

import "errors"

func GetPost(typePost string) (IPost, error) {
	if typePost == "news" {
		return newNewsPost(), nil
	}
	if typePost == "normal" {
		return newNormalPost(), nil
	}
	return nil, errors.New("typePost incorrect")
}