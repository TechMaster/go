package repo

import (
	"one_many/model"

	"github.com/brianvoe/gofakeit/v6"
)

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
