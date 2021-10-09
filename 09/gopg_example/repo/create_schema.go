package repo

import (
	"gopgdemo/model"

	"github.com/go-pg/pg/v10/orm"
)

func CreateSchema() error {
	orm.RegisterTable(new(model.Track_Course))
	
	models := []interface{}{
		
		new(model.Track),
		new(model.Course),
		new(model.Track_Course),
		new(model.Writer),
		new(model.Story),
	}

	for _, table := range models {
		err := DB.Model(table).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
