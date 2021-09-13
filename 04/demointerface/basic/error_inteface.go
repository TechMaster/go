package basic

import (
	"fmt"
	"os"
)

type OneMountError struct {
	Msg      string
	Division string
	Module   string
	Func     string
	Line     int
}

type NetworkError struct {
	Msg    string
	IP     string
	Module string
	Func   string
}

func (err OneMountError) Error() string {
	return fmt.Sprintf("OneMount Error: %s at \n  {Division: %s, Module: %s, Func: %s, Line: %d}", err.Msg, err.Division, err.Module, err.Func, err.Line)
}

func (err NetworkError) Error() string {
	return fmt.Sprintf("Network Error: %s at \n  {IP: %s, Module: %s, Func: %s}", err.Msg, err.IP, err.Module, err.Func)
}

func openMissingFile() error {
	file, err := os.OpenFile("invalid path", os.O_RDWR, 0644)

	if err != nil {
		return OneMountError{Msg: err.Error(), Division: "OneHousing", Module: "CRM", Func: "openMissingFile", Line: 32}
	}
	defer file.Close()
	return nil
}

func connectMessageQueue() error {
	return NetworkError{
		Msg:    "Service is not available",
		IP:     "192.168.1.1",
		Module: "Message Queue",
		Func:   "connectMessageQueue",
	}
}

func DemoCustomError() {
	if err := openMissingFile(); err != nil {
		fmt.Println(err)
	}

	if err := connectMessageQueue(); err != nil {
		fmt.Println(err)
	}
}
