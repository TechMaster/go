# Báo lỗi bao gồm Stack Trace

Khi trả về lỗi cần có đủ thông tin :

1. Hàm gây lỗi
2. Dòng gây lỗi
3. Ở file nào?

Thường các hàm gọi liên tiếp nhau, thì lỗi cần hiển thị hàm trong cùng (trên đỉnh stack), rồi đến các hàm kế tiếp để dễ dàng dò lỗi.

Để làm được việc này cần 
```go
runtime.Caller(calldepth)
runtime.FuncForPC(pc)
```

`runtime.Caller(calldepth)`
tham số calldepth: bỏ qua N hàm bên trên call stack. Ví dụ N = 2 thì bỏ qua
- getFrame chính là hàm hiện tại, được AppendStackTrace gọi
- AppendStackTrace: hàm này không cần in ra vì nó phục vụ add stack lỗi thôi
và chỉ lấy từ hàm gọi AppendStackTrack
trả về :
- pc: program  counter
- file: file chứa hàm
- line: dòng hàm đang thực thi
- ok: true nếu lấy thành công

Kết quả báo lỗi như sau:
```
/Volumes/CODE/go/14/stacktrace/main.go:9 main.main
testOne: /Volumes/CODE/go/14/stacktrace/main.go:18 main.testOne
testTwo: /Volumes/CODE/go/14/stacktrace/main.go:23 main.testTwo
/Volumes/CODE/go/14/stacktrace/main.go:28 main.testThree
internal error
````