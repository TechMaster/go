package repo

import (
	"one_many/model"

	"github.com/brianvoe/gofakeit/v6"
)

func Get_Foo_Bar() (foos []model.Foo, err error) {
	err = DB.Model(&foos).Relation("Bars").Select()
	if err != nil {
		return nil, err
	}
	return foos, nil
}

func Get_Foo_By_ID(id string) (foo *model.Foo, err error) {
	foo = new(model.Foo)
	err = DB.Model(foo).Relation("Bars").Where("id = ?", id).Limit(1).Select()
	//err = DB.Model(foo).Relation("Bars").Where("id = ?", id).First() --> WHERE (id = 'ox-01') ORDER BY "foo"."id" LIMIT 1

	if err != nil {
		return nil, err
	}
	return foo, nil
}

func Get_Bar_By_ID(id string) (bar *model.Bar, err error) {
	bar = &model.Bar{
		Id: "bar3",
	}

	err = DB.Model(bar).Relation("Foo").WherePK().Select()

	if err != nil {
		return nil, err
	}
	return bar, nil
}

/*
Tạo nhiều foo và bar. Chú ý:
- Cần sử dụng transaction để rollback khi có lỗi
- go-pg không lưu được relation bars nằm trong foo mà phải tạo tách biệt foo sau đó rồi bar
*/
func Create_Foo_Bar() (err error) {
	transaction, err := DB.Begin()
	if err != nil {
		return err
	}
	for i := 0; i < 5; i++ {
		foo_id := NewID()
		_, err = transaction.Model(&model.Foo{
			Id:   foo_id,
			Name: gofakeit.Animal(),
		}).Insert()

		if !check_err(err, transaction) {
			return err
		}

		for j := 0; j < 2+random.Intn(2); j++ {
			_, err = transaction.Model(&model.Bar{
				Id:    NewID(),
				Name:  gofakeit.Animal(),
				FooId: foo_id,
			}).Insert()
			if !check_err(err, transaction) {
				return err
			}
		}
	}
	return transaction.Commit()
}
