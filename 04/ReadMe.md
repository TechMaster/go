# Chủ đề hôm nay về Interfance

Các bạn hay chạy code ở thư mục [demointerface](demointerface) trước.

Sau đó chúng sẽ xây dựng chức năng gửi email có nhiều biến thể khác nhau ở [mailer](mailer)

## Tuân thủ interface - implement interface

- Chỉ làm việc trên các prototype function mà đối tượng tuân thủ Interface
- Không bận tâm đến kiểu thực sự của đối tượng
- Tính đa hình giúp lập trình viên quản lý được một nhóm đối tượng rất đa dạng
- Một đối tượng trong Golang có thể tuân thủ nhiều interface
- Để một đối tượng được gọi là `tuân thủ` một interface thì đối tượng đó phải implement tất cả các prototype function trong interface đó.
- Việc tuân thủ interface trong Golang là `implicit` không cần khai báo, cứ tuân thủ đủ các prototype function của một interface là đạt
```go
shapes := []Shape{rec, cir, tri}
for _, shape := range shapes {
  //Polymorphism
  fmt.Printf("%s perimeter = %.2f\n", shape, shape.Perimeter())
}
```
## Empty Interface `interface{}`

Kiểu Empty Interface chứa đối tượng thuộc bất kỳ kiểu nào bởi vì Empty Interface không chứa bất kỳ prototype function nào.

- Empty Interface là `interface{}`
- `interface{}` chấp nhận bất kỳ kiểu gì int, string, array, slice, struct, func
- Dùng `interface{}` khi nào? dùng khi không biết trước kiểu sẽ truyền vào

## Kiểm tra kiểu Interface
Cú pháp kiểm tra kiểu bằng if assigment

```go
if f, ok := i.(string); ok { //Golang if assignment statement
  fmt.Println(f)
}
```

Kiểm tra kiểu bằng cú pháp `switch`
```go
for _, v := range arrayAnyType {
  switch v.(type) {
  case string:
    fmt.Printf("%v is string\n", v)
  case Person:
    fmt.Printf("%v is Person\n", v)
  case func(int, int) int:
    fmt.Printf("%v is func(int, int) int\n", v)
  }

}
```
> Gợi ý: khi lưu trữ thì bao dung (dùng kiểu `interface{}`) nhưng khi làm việc với từng đối tượng cụ thể, cần phải tìm hiểu rõ đối tượng đó là kiểu nào hoặc tuân thủ interface nào. Nếu dùng sai, lập tức sẽ gặp lỗi panic

## Alias 

Trong Golang chúng ta có thể định nghĩa alias của một kiểu xác định. Mục tiêu là đặt tên gợi nhớ dễ hiểu. Ví dụ
```go
type Operator2operand = func(int, int) int
type AnyType = interface{}
```

## Receiver không chấp nhận một số kiểu căn bản
```go
func (p int) String() string {
	return "it is basic int type"
}
```
Compiler báo lỗi `invalid receiver int (basic or unnamed type)`
Lỗi tương tự
```go
func (p interface{}) String() string {
	return "it is basic int type"
}
```
Chúng ta phải chuyển đối tượng đó thành tham số đầu tiên truyền vào hàm (không phải receiver)
```go
func SpellNumber(i int) string {
	return "it is basic int type"
}
```

## Xử lý lỗi trong Golang
Golang không có try catch exception nó chỉ có duy nhất một interface

```go
error interface {
  Error() string
}
```

Lập trình thoải mái tuỳ biến bằng cách tạo mới các cấu trúc lỗi khác nhau.

Việc định nghĩa Error theo các bước sau đây:
1. Theo business domain ~ division
2. Theo application architect ~ module > function > line
3. Theo error level: cấp độ nghiêm trọng

Việc xử lý lỗi cũng có các cấp độ như sau:
1. Xử lý lỗi trong lúc phát triển, debug chỉ cần xuất ra console.
2. Xử lý lỗi ở monolithic app chỉ cần xuất ra file
3. Xử lý lỗi ở hệ thống microservice cần xử dụng Centralize Logging pattern. Tuy nhiên pattern có rủi ro Single Point Failure. Nó mà chết thì lấy ai xử lý lỗi !
Do đó cần có thư viện log lỗi đầu tiên thử kết nối đến Centralize Logging nếu kết nối thành công thì push lỗi vào đó, nếu không kết nối được thì phải ghi ra file.

Centralize Logging pattern nếu sử dụng Event Streaming kiểu như Kafka hoặc rẻ tiền dùng Redis Stream.
Sẽ có một dịch vụ lắng nghe Error Stream rồi phân phối báo lỗi đến nơi cần thiết.

> Báo lỗi và Xử lý lỗi là một nghệ thuật. Ai không làm chủ được nghệ thuật này thì còn mệt mỏi quản trị hệ thống dài dài.

## Gửi email

Có 101 cách gửi email từ những cách gửi cực đơn giản mở SMTP connection rồi gửi cho đến sử dụng OAuth2, Email Service....

Do đó cần phải xây dựng một interface gửi email thống nhất
```go
type MailSender interface {
	SendPlainEmail(to []string, subject string, body string) error
	SendHTMLEmail(to []string, subject string, data map[string]interface{}, layout ...string) error
}
```

Trong quá trình phát triển sản phẩm: trời
Prototype > Debug > Monolithic App > Small Scale Microservice > Large Scale Microservice > Global Scale Microservice

Chúng ta sẽ phát triển nhiều biến thể gửi email. Nếu trong hệ thống có 1000 dịch vụ, mỗi dịch vụ có 4 điểm gọi hàm gửi email.

Việc định nghĩa tốt (rõ ràng, bao được tất cả trường hợp) interface gửi email cực kỳ quan trọng. Nó giúp chúng ta thoải mái bổ xung các biến thể gửi email, mà không phải sửa đổi code ở 1000 x 4 điểm. Sửa đổi code chưa khổ đau bằng việc phải re-deploy 1000 dịch vụ.

> Do đó luôn ghi nhớ rằng, hãy lên kế hoạch, nhìn xa hơn về tương lai để định nghĩa interface. Luôn luôn dùng interface trong Golang để tạo ra các biến thể (implementations). Việc này đảm bảo cấu hình ở một điểm, không phải sửa code ở tất cả các điểm còn lại.


Bài hôm nay về Interface kết thúc được rồi đó.
Cảm ơn các bạn đã lắng nghe.



## Chuyển đổi map[string]interface{} sang struct

## Tham khảo về Interface
- [Golang Tour Interfaces](https://tour.golang.org/methods/9)
- [Empty Interface](https://tour.golang.org/methods/14)
- [Type assertion / Type switch](https://golangbot.com/interfaces-part-1/)
- [func-method-interface](https://github.com/zalopay-oss/go-advanced/blob/master/ch1-basic/ch1-04-func-method-interface.md)