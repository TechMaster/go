package bar

import (
	"github.com/rotisserie/eris"
)

func Connect_to_db() error {
	return eris.New("Failed to connect DB")
}

func Call_other_rest() error {
	return eris.New("connect API failed")
}
