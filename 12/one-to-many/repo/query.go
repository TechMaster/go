package repo

import (
	"one-many/model"
	"one-many/util"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func GetFooBar(db *gorm.DB) (foos []model.Foo, err error) {
	if err := db.Preload("Bars").Find(&foos).Error; err != nil {
		return nil, err
	}
	return foos, nil
}

func GetFooById(db *gorm.DB, id string) (foo model.Foo, err error) {
	if err := DB.Preload("Bars").Find(&foo, "foo.id = ?", id).Error; err != nil {
		return model.Foo{}, err
	}

	return foo, nil
}

func GetBarById(db *gorm.DB, id string) (bar model.Bar, err error) {
	bar = model.Bar{
		Id: id,
	}

	if err = db.Preload("Foo").Find(&bar).Error; err != nil {
		return model.Bar{}, err
	}

	return bar, nil
}

func CreateData(db *gorm.DB) (err error) {
	var foos []model.Foo

	for i := 0; i < 5; i++ {
		foo_id := util.NewID()
		foo := model.Foo{
			Id:   foo_id,
			Name: gofakeit.Animal(),
		}

		for j := 0; j < 2+random.Intn(2); j++ {
			bar := model.Bar{
				Id:    util.NewID(),
				Name:  gofakeit.Animal(),
				FooId: foo_id,
			}

			foo.Bars = append(foo.Bars, bar)
		}

		foos = append(foos, foo)
	}

	if err := db.Create(&foos).Error; err != nil {
		return err
	}

	return nil
}
