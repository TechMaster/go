package bar

import (
	"github.com/TechMaster/eris"
)

func Bar() error {
	return eris.New("Cannot connect to DB").InternalServerError().SetData(map[string]interface{}{
		"host": "localhost",
		"port": "5432",
		"db":   "inventory",
	}).SetType(eris.SYSERROR)
}
