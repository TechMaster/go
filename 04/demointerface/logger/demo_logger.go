package logger

import "demointerface/basic"

func DemoLogger() {
	var logger Logger
	logger = OneMountLogger{} // sau đó có thể đổi thành TechMasterLogger

	err := basic.OneMountError{
		Msg:      "Failed to register account",
		Division: "VinID",
		Module:   "Online Shopping",
		Func:     "Account_Register",
		Line:     120,
	}
	logger.Log(err)
}
