package repo

import (
	"one-to-one/model"
	"one-to-one/util"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func GetAllPerson(db *gorm.DB) (persons []model.Person, err error) {
	if err := db.Preload("Card").Find(&persons).Error; err != nil {
		return nil, err
	}
	return persons, nil
}

func GetPersonById(db *gorm.DB, id string) (person *model.Person, err error) {
	if err := db.Preload("Card").Find(&person, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return person, nil
}

func CreateData(db *gorm.DB) (err error) {
	var persons []model.Person
	for i := 0; i < 5; i++ {

		personId := util.NewID()

		person := model.Person{
			Id:       personId,
			FullName: gofakeit.Animal(),
			Email:    gofakeit.Email(),
			Card: model.Card{
				Id:       util.NewID(),
				Seri:     util.NewID(),
				PersonId: personId,
			},
		}

		persons = append(persons, person)
	}

	if err := db.Create(&persons).Error; err != nil {
		return err
	}

	return nil
}
