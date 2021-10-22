package bar

import (
	"demoeris/rock"

	"github.com/rotisserie/eris"
)

func Connect_to_db() error {
	//return eris.New("Failed to connect DB")
	if err := rock.ConnectRedis(); err != nil {
		return eris.Wrap(err, "Lỗi kết nối Redis")
	} else {
		return nil
	}

	//return errors.New("Failed to connect")
}

func Call_other_rest() error {
	return eris.New("connect API failed")
}
