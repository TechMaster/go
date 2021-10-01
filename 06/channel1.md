> Series các bài viết này tập trung về một chủ đề Channels trong Golang. Golang có một đặc sản là concurrency programming (các tác vụ chạy đồng thời) sử dụng go routine nhỏ gọn hơn thread. Để chia sẻ dữ liệu giữa các go routine, cách đầu tiên mà ai cũng nghĩ đến là share memory : các go routine cùng truy xuất đến biến chung. Để đảm bảo tính toàn vẹn của biến chung khi nhiều go routine cùng thay đổi, người ta sẽ sử dụng kỹ thuật đồng bộ synchronization, mutex lock. Tuy nhiên Golang đưa ra một kỹ thuật mới hay hơn đó là Channel. Rob Pike tác giả của Golang đúc kết như sau : _concurrent programming is don't (let computations) communicate by sharing memory, (let them) share memory by communicating (through channels)._. Tạm dịch là lập trình đồng thời không truyền tin bằng dùng chung bộ nhớ mà dùng chung bộ nhớ bằng cách truyền tin qua kênh (channels). [Toàn bộ code minh hoạ cho series này các bạn có thể tải về ở đây](https://github.com/TechMaster/golang/blob/main/13Channel/01Basic/ChannelString.go)

## 1. Channels mô phỏng máy bơm nước, đường ống nước, van khoá

Channel dịch sát ra tiếng Việt là kênh, nó giống như một đường ống dẫn nước nối giữa một bên là máy bơm nước và bên đường ống là nơi nhận nước. Golang channel hoạt động như sau: chúng ta đẩy thông tin vào channel (bơm nước), ở đâu bên kia cần phải mở van khoá để nhận nước. Nếu không mở van cho nước lưu thông, máy bơm sẽ bị cháy tương tự ứng dụng sẽ bị lỗi.

![](https://techmaster.vn/media/static/36/c32ohdc51cofhqf7v88g)

Xem file ChannelString.go:

```go
package main
func DemoChannelString() {
    pipe := make(chan string) //tạo một kênh kiểu string
    pipe <- "water" //bơm một lượng nước
}
```

Đoạn code này gây lỗi `fatal error: all goroutines are asleep - deadlock!`. Tại sao vậy? vì đầu kia đường ống, không có nơi tiếp nhận, van nước đóng, áp suất tăng nhanh chóng khiến máy bơm không thể bơm.

Tôi sửa lại code như dưới.

```go
func DemoChannelString() {
    pipe := make(chan string) //tạo một kênh kiểu string
    pipe <- "water" //bơm "water" vào đường ống
    receiver := <-pipe //nhận dữ liệu từ đường ống đổ vào biến receiver
    fmt.Println(receiver)  //in biến receiver
}
```

Chạy lại thì chương trình vẫn đổ vỡ cùng lỗi `fatal error: all goroutines are asleep - deadlock!`. Tại sao vậy?

Trả lời: vì đây là cách viết tuần tự chạy từng dòng lệnh: một anh thợ bơm dữ liệu vào đường ống `pipe` thì đầu bên kia thực ra chưa mở van. Kết quả là bơm lập tức cháy máy, ứng dụng lỗi luôn ở dòng `pipe <- "water"` khi chưa kịp mở van khoá đầu bên kia đường ống.

### 1.1 Sử dụng go routine để khắc phục

Để khắc phục tình trạng này, anh thợ gọi thêm một người bạn. Họ thống nhất cùng hành động (concurrent coordianate): khi nào tôi bật máy bơm, anh hãy mở van để nhận nước nhé.

![](https://techmaster.vn/media/static/36/c32oqcc51cofhqf7v890)

```go
func DemoChannelString() {
    pipe := make(chan string)
    go func() {
        pipe <- "water"
    }()
    receiver := <-pipe
    fmt.Println(receiver)
}
```

Hãy chú ý đến đoạn lệnh. Khi thực thi trong `go func()` lệnh bơm nước `pipe <- "water"` không ngăn lệnh tiếp theo thực thi. Điều này giống với 2 anh thợ đồng thời phối hợp bơm và nhận ở hai go routing độc lập, không phải chờ đợi nhau. Giờ bạn đã hiểu `go routine` hiện thực hoá concurrency programming rồi đó.

```go
go func() {
    pipe <- "water"
}()
```

### 1.2 Sử dụng channel buffer để khắc phục

Ở cách đầu tiên, mỗi lần bơm nước, chúng ta cần đến 2 anh thợ (2 go routine) phối hợp bơm và mở van khoá. Còn ở cách này chúng ta chỉ cần một anh thợ nhưng phải dùng đến bể trung gian để máy bơm không bị quá áp suất khi van đầu kia khoá.

![](https://techmaster.vn/media/static/36/c32p2nc51cofhqf7v89g)

Duy nhất một anh thợ sẽ thực hiện các bước:

1. Thay đường ống bình thường bằng đường ống có bể chứa phụ
2. Bật máy bơm. Nhờ bể chứa phụ nên dù đầu bên kia đường ống, van đang khoá, nhưng máy bơm không bị quá áp suất
3. Chạy sang đầu bên kia mở van. Nước được bơm sang đầu bên kia thành công

```go
func DemoChannelString() {
    pipe := make(chan string, 1)
    pipe <- "water"
    receiver := <-pipe
    fmt.Println(receiver)
}
```

Kết quả in ra

```
water
```

### 1.3 Khi buffer bị quá tải

Anh thợ bơm phấn khích với cải tiến của mình, quyết định tăng gấp đôi lưu lượng khi mở máy bơm

```go
func DemoChannelString() {
    pipe := make(chan string, 1)
    pipe <- "water"
    pipe <- "water"  //Bơm gấp đôi lượng nước mà bể phụ có thể chứa
    receiver := <-pipe
    fmt.Println(receiver)
}
```

Kết quả ứng dụng lại đổ vỡ vì lượng nước lớn hơn bể phụ có thể chứa, áp suất tăng đột biến, máy bơm cháy. Anh ta sửa chữa bằng cách tăng gấp đôi thể tích bể phụ

```go
func DemoChannelString() {
    pipe := make(chan string, 2)
    pipe <- "water"
    pipe <- "water"
    receiver := <-pipe
    fmt.Println(receiver)
    receiver = <-pipe
    fmt.Println(receiver)
}
```

kết quả in ra

```
ping
ping
```

### 1.4 Khi nào cần phải đóng máy bơm

Thay vì phải lặp lại quá nhiều lệnh

```go
receiver := <-pipe
fmt.Println(receiver)
receiver = <-pipe
fmt.Println(receiver)
```

Tôi cải tiến lại bằng cách sử dụng vòng lặp for liên tục đọc ra từ

```go
pipe := make(chan string, 2)
pipe <- "water"
pipe <- "water"

for receiver := range pipe {
    fmt.Println(receiver)
}
```

Sau khi in ra

```
water
water
```

thì ứng dụng lại bị đổ vỡ `fatal error: all goroutines are asleep - deadlock!`

Hiện tượng này giống như đầu bên kia máy bơm là một bình đun nước thô sơ. Khi van mở, nước tiếp vào thì nó hoạt động trơn tru. Nhưng khi van mở, nước không vào, bình cạn bị đun nóng dẫn bình bị nung chảy. Do đó khi không bơm nước nữa, chúng ta cần đóng van đường ống: `close(pipe)`

```go
pipe := make(chan string, 2)
pipe <- "water"
pipe <- "water"
close(pipe)  //Dừng bơm, khoá van
for receiver := range pipe {
    fmt.Println(receiver)
}
```

Mọi thứ lại chạy ổn.

### 1.5 Cách viết khác gói phần bơm dữ liệu vào channel trong `go func`

Hỏi : Ý nghĩa của channel để làm gì?

Đáp: Channel là để các tác vụ chạy ở các go routine khác nhau, có thể chia sẻ dữ liệu bằng cách truyền tin (bơm dữ liệu qua channel).

Do đó đoạn code dưới đây mới thực sự là có ý nghĩa vì phần bơm dữ liệu vào channel được chạy trong một go routine khác.

```go
pipe := make(chan string, 2)
go func() {
    pipe <- "water"
    pipe <- "water"
    close(pipe)
}()

for receiver := range pipe {
    fmt.Println(receiver)
}
```

## 2. Kết luận

Ở bài đầu tiên, chúng ta làm quen với khái niệm buffer qua ví dụ thực tế:

- channel <--> đường ống
- go routine <--> hai anh thợ cùng phối hợp bơm - nhận nước
- channel buffer <--> đường ống có bể phụ
- Lỗi `fatal error: all goroutines are asleep - deadlock!` xảy ra ở 2 trường hợp:
    - Nước bơm, van khoá: đường ống bị quá tải
    - Nước không bơm, van không khoá, đầu kia vẫn nhận --> không có nước, cháy bình đun
- close(channel) <--> đóng đường ống không bơm nước nữa

Hôm rồi có một bài viết của tiến sỹ Ái Việt về tình trạng lập trình Việt nam chỉ mải code mà không chịu hiểu bản chất vật lý. Tiến sỹ Ái Việt nhận một đống gạch đá đủ xây mấy căn biệt thự. Quá oan cho ông Ái Việt. Thực tế là để hiểu rõ bản chất của các kỹ thuật trong lập trình đa nhiệm, đồng thời, nếu các bạn nắm vững các nguyên tắc vật lý, các bạn sẽ thấy dễ hiểu vô cùng. Còn ngược lại, các bạn sẽ chỉ dùng một cách máy móc.

## Tham khảo

- [https://go101.org/article/channel.html](https://go101.org/article/channel.html)
- [Understand Golang channels and how to monitor with Grafana (Part 1/2)](https://dev.to/ahmedash95/understand-gochannels-and-how-to-monitor-with-grafana-154)