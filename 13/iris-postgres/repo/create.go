package repo

import (
	"iris-postgres/model"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

func MockData() (err error) {
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		userId := NewID()
		_, err = transaction.Model(&model.User{
			Id:        userId,
			FullName:  gofakeit.Animal(),
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
			CreatedAt: gofakeit.Date(),
		}).Insert()

		if err != nil {
			transaction.Rollback()
			return err
		}

		for j := 0; j < 5+rand.Intn(5); j++ {
			_, err = transaction.Model(&model.Post{
				Id:        NewID(),
				Title:     gofakeit.Quote(),
				Content:   gofakeit.LoremIpsumSentence(200),
				AuthorId:  userId,
				CreatedAt: gofakeit.Date(),
			}).Insert()

			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	return transaction.Commit()
}
