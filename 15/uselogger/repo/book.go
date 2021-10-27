package repo

import (
	"uselogger/model"

	"github.com/TechMaster/eris"
)

func Get_books() (books []model.Book, err error) {
	return nil, eris.New("Cannot connect to database").SetData(map[string]interface{}{
		"host": "localhost",
		"port": "5432",
	}).SetType(eris.SYSERROR)
}
