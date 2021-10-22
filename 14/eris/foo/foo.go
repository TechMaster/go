package foo

import (
	"demoeris/bar"
)

func Foo() error {
	return bar.Connect_to_db()
}
