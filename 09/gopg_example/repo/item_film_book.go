package repo

import (
	"gopgdemo/model"

	"github.com/brianvoe/gofakeit/v6"
)

func Create_films_books() (err error) {
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		//Tạo Book
		_, err = transaction.Model(&model.Book{
			Item: model.Item{
				Id:    NewID(),
				Name:  gofakeit.Name(),
				Price: random.Intn(300),
			},
			Pages: random.Intn(500),
		}).Table("test.book").Insert()
		if !check_err(err, transaction) {
			return err
		}
		//Tạo Film
		filmID := NewID()
		_, err = transaction.Model(&model.Film{
			Item: model.Item{
				Id:    filmID,
				Name:  gofakeit.Name(),
				Price: random.Intn(300),
			},
			Director: gofakeit.Dog(),
		}).Table("test.film").Insert()
		if !check_err(err, transaction) {
			return err
		}
		//Thêm Actor cho phim
		for j := 0; j < random.Intn(6); i++ {
			_, err = transaction.Model(&model.Actor{
				Id:     NewID(),
				Name:   gofakeit.Animal(),
				FilmId: filmID,
			}).Insert()
			if !check_err(err, transaction) {
				return err
			}
		}
	}
	return transaction.Commit()

}
