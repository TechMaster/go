package repo

import (
	"gopgdemo/model"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func Create_product_price() (err error) {
	transaction, err := DB.Begin()

	for i := 0; i < 20; i++ {
		cpu := "Intel"
		if random.Intn(3) == 1 {
			cpu = "AMD"
		}
		id := NewID()
		_, err = transaction.Model(&model.Laptop{
			Id:   id,
			Name: gofakeit.Name(),
			Cpu:  cpu,
		}).Insert()

		if !check_err(err, transaction) {
			return err
		}

		//tạo giá thay đổi
		for j := 2; j < 3+random.Intn(4); j++ {
			_, err = transaction.Model(&model.PriceHistory{
				Id:          NewID(),
				ItemId:      id,
				Type:        "laptop",
				Price:       100 + random.Intn(1000),
				CreatedDate: time.Now().Add(time.Duration(-random.Intn(100)) * 24 * time.Hour),
			}).Insert()
			if !check_err(err, transaction) {
				return err
			}
		}
	}

	for i := 0; i < 30; i++ {

		id := NewID()
		_, err = transaction.Model(&model.Bike{
			Id:     id,
			Name:   gofakeit.Animal(),
			Volume: 50 + random.Intn(500),
		}).Insert()
		if !check_err(err, transaction) {
			return err
		}

		//tạo giá thay đổi
		for j := 2; j < 3+random.Intn(4); j++ {
			_, err = transaction.Model(&model.PriceHistory{
				Id:          NewID(),
				ItemId:      id,
				Type:        "bike",
				Price:       1400 + random.Intn(20000),
				CreatedDate: time.Now().Add(time.Duration(-random.Intn(100)) * 24 * time.Hour),
			}).Insert()
			if !check_err(err, transaction) {
				return err
			}
		}
	}
	return transaction.Commit()
}

/*
select * from price_history ph where item_id ='_PflAeeW' order by created_date  desc limit 1
*/
func Get_latest_price_of_item(name string) (price int) {
	return 0
}
