package error_stack

//https://gist.github.com/vardius/56b224de0a69522a24f021642449db17
import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
)

/*
Định nghĩa một struct AppError mới
*/
type AppError struct {
	trace string //Lưu calling stack trace
	err   error  //thực sự lưu lỗi
}

// Nó tuân thủ interface của Error vậy có thể trả về nhưng error bình thường
func (e *AppError) Error() string {
	return fmt.Sprintf("%s\n%s", e.trace, e.err)
}

//Trả về thuộc tính err bên trong struct
func (e *AppError) Unwrap() error {
	return e.err
}

// StackTrace returns the string representation of the error stack trace,
// includeTrace appends caller pcs frames to each error message if possible.
func (e *AppError) StackTrace() (string, error) {
	var buf bytes.Buffer

	if e.trace != "" {
		if _, err := fmt.Fprintf(&buf, "%s", e.trace); err != nil {
			return "", err
		}
	}

	if e.err == nil {
		return buf.String(), nil
	}

	var next *AppError
	if errors.As(e.err, &next) {
		stackTrace, err := next.StackTrace()
		if err != nil {
			return "", err
		}

		buf.WriteString(fmt.Sprintf("\n%s", stackTrace))
	} else {
		return fmt.Sprintf("%s\n%s", buf.String(), e.err), nil
	}

	return buf.String(), nil
}

/*
Thêm lỗi đồng thời bổ xung tên hàm, tên file, dòng số bao nhiêu vào biến trace
*/
func AppendStackTrace(err error) *AppError {
	if err == nil {
		panic("nil error provided")
	}

	var buf bytes.Buffer

	frame := getFrame(0) //Bỏ qua hàm AppendStackTrace và getFrame

	fmt.Fprintf(&buf, "%s", frame.File)
	fmt.Fprintf(&buf, ":%d", frame.Line)
	fmt.Fprintf(&buf, " %s", frame.Function)

	return &AppError{
		err:   err,
		trace: buf.String(),
	}
}

/*
Lấy thông tin: tên hàm gây lỗi, tên file chưa hàm và dòng gây lỗi
*/
func getFrame(calldepth int) *runtime.Frame {
	/*
		runtime.Caller(calldepth)
		tham số calldepth: bỏ qua N hàm bên trên call stack. Ví dụ N = 2 thì bỏ qua
		- getFrame chính là hàm hiện tại, được AppendStackTrace gọi
		- AppendStackTrace: hàm này không cần in ra vì nó phục vụ add stack lỗi thôi
		và chỉ lấy từ hàm gọi AppendStackTrack
		trả về :
		- pc: program  counter
		- file: file chứa hàm
		- line: dòng hàm đang thực thi
		- ok: true nếu lấy thành công
	*/

	pc, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		return nil
	}

	frame := &runtime.Frame{
		PC:   pc,
		File: file,
		Line: line,
	}

	funcForPc := runtime.FuncForPC(pc)
	if funcForPc != nil {
		frame.Func = funcForPc
		frame.Function = funcForPc.Name() //Lấy tên  Function
		frame.Entry = funcForPc.Entry()
	}

	return frame
}
