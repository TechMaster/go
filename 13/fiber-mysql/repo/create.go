package repo

import (
	"fiber-postgres/model"
	"math/rand"

	"github.com/brianvoe/gofakeit/v6"
)

func MockData() (err error) {
	transaction := DB.Begin()
	if err := transaction.Error; err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		userId := NewID()
		err = transaction.Create(&model.User{
			Id:        userId,
			FullName:  gofakeit.Animal(),
			Email:     gofakeit.Email(),
			Phone:     gofakeit.Phone(),
			CreatedAt: gofakeit.Date(),
		}).Error

		if err != nil {
			transaction.Rollback()
			return err
		}

		for j := 0; j < 5+rand.Intn(5); j++ {
			err = transaction.Create(&model.Post{
				Id:        NewID(),
				Title:     gofakeit.Quote(),
				Content:   gofakeit.LoremIpsumSentence(200),
				AuthorId:  userId,
				CreatedAt: gofakeit.Date(),
			}).Error

			if err != nil {
				transaction.Rollback()
				return err
			}
		}
	}

	return transaction.Commit().Error
}
