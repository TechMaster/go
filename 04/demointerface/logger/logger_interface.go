package logger

import (
	"demointerface/basic"
	"fmt"
)

type Logger interface {
	Log(err error)
}

type OneMountLogger struct{}
type TechMasterLogger struct{}

func (logger OneMountLogger) Log(err error) {
	switch err.(type) {
	case basic.OneMountError:
		fmt.Println("Send to One Mount QA Team")
		fmt.Println(err)
	case basic.NetworkError:
		fmt.Println("Send to DevOps Team")
		fmt.Println(err)
	default:
		//Write to log file
		fmt.Println(err)
	}
}

func (logger TechMasterLogger) Log(err error) {
	fmt.Println(err)
}
