package bar

import (
	"errors"
	"fmt"
)

func Bar() error {
	defer func() {
		fmt.Println("Clean up resource at the end of Bar")
	}()
	return errors.New("Error from Bar function")
}
