# Eris Error

Bản chất vẫn là để lấy ra được Calling Stack hàm gây lỗi --> các hàm gọi nó
dùng runtime.Caller(depth int)

Muốn dùng được eris track strace thì phải tạo lỗi kiểu Eris


Điểm chưa ưng lắm của eris error là:

Cấu trúc lỗi dạng private, nên không customize (chế cháo) theo ý đồ của mình
```go
type rootError struct {
	global bool   // flag indicating whether the error was declared globally
	msg    string // root error message
	ext    error  // error type for wrapping external errors
	stack  *stack // root error stack trace
}
```

Tôi muốn không chỉ in ra strack track, mà còn:

1. Mức độ nghiêm trọng của lỗi
2. Phân loại lỗi theo http status code
3. Thêm vào data để trả về cho client

Development Phase - Production Phase cần báo lỗi chi tiết mức độ khác nhau.

