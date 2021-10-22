package main

import (
	"fmt"
	"techeris/foo"

	"github.com/TechMaster/eris"
)

var ErisStringFormat eris.StringFormat

func main() {
	ErisStringFormat = eris.StringFormat{
		Options: eris.FormatOptions{
			InvertOutput: false, // flag that inverts the error output (wrap errors shown first)
			WithTrace:    true,  // flag that enables stack trace output
			InvertTrace:  true,  // flag that inverts the stack trace output (top of call stack shown first)
			WithExternal: false,
			Top:          3, // Chỉ lấy 3 dòng lệnh đầu tiên
			//Mục tiêu để báo lỗi gọn hơn, stack trace đủ ngắn
		},
		MsgStackSep:  "\n",  // separator between error messages and stack frame data
		PreStackSep:  "\t",  // separator at the beginning of each stack frame
		StackElemSep: " | ", // separator between elements of each stack frame
		ErrorSep:     "\n",  // separator between each error in the chain
	}
	if err := foo.Foo(); err != nil {
		switch e := err.(type) {
		case *eris.Error: //Lỗi kiểu eris
			logErisError(e)
		default: //Lỗi thông thường
			fmt.Println(err.Error()) //In ra console
		}
	}
}

//Hàm chuyên xử lý Eris Error có Stack Trace. Chỉ áp dụng với cấp độ lỗi ERROR, SYSERROR, PANIC
func logErisError(err *eris.Error) {
	formattedStr := eris.ToCustomString(err, ErisStringFormat) //Định dạng lỗi Eris

	//Chỗ này log ra console
	if err.ErrType > eris.ERROR { //Với lỗi cao hơn ERROR gồm SYSERROR và PANIC, in ra mầu đỏ và ghi ra file
		colorReset := string("\033[0m")
		colorMagenta := string("\033[35m")
		fmt.Println(colorMagenta, formattedStr, colorReset)
	} else {
		fmt.Println(formattedStr) //Error Level
	}
}
