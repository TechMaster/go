package demopanic

import (
	"errors"
	"fmt"
)

func Connect_to_db() {
	defer func() {
		fmt.Println("Clean up db connection")
	}()
	panic(errors.New("failed to connect to DB"))
}
