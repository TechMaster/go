> Đây là bài số 2, tiếp theo [bài 01](https://techmaster.vn/posts/36560/tat-tat-ve-channel-trong-gobai-01). Trong bài này chúng ta quan sát hiện tượng hai go routine chạy đồng thời, truyền và nhận tin. [Code demo chi tiết ở đây, các bạn nên tải về chạy thử rồi tự trải nghiệm nhé](https://github.com/TechMaster/golang/blob/main/13Channel/01Basic/DemoChannel.go)

## 1. Go routine chưa kịp nhận dữ liệu, hàm đã thoát

Hãy chạy thử hàm dưới đây

```go
func DemoChannel0() {
    pipe := make(chan string)
    go func() { //1
        for receiver := range pipe { //2
            fmt.Println(receiver) //3
        }
    }()
    pipe <- "water 1" //4
    close(pipe)
}
```

Màn hình console chả in ra gì cả. Khi tôi đặt break point ở điểm `//1, 2, 3, 4` thì thấy rằng ứng dụng chạy theo thứ tự `1 -> 4 -> 2 -> 3`. Dữ liệu được bơm vào channel, nhưng ngay sau đó ứng dụng thoát khiến cho phần receiver không kịp nhận dữ liệu hoặc không kịp in ra màn hình.

![](https://techmaster.vn/media/static/36/c32torc51cofhqf7v8ag)

Khi tôi thử bơm liên tiếp vài đoạn dữ liệu

```go
pipe := make(chan string)
go func() {
    for receiver := range pipe {
        fmt.Println(receiver)
    }
}()
pipe <- "water 1"
pipe <- "water 2"
pipe <- "water 3"
close(pipe)
```

Kết quả in ra màn hình bị sót mất `water 3`

```
water 1
water 2
```

Bổ xung một lệnh delay để trễ lại việc hàm thoát

```go
close(pipe)
time.Sleep(time.Millisecond)
```

Kết quả in ra đầy đủ

```
water 1
water 2
water 3
```

Đây là một cách lập trình lang băm. Trễ bao nhiêu thì đủ? Để chắc ăn tôi tạo trễ hẳn 1 giây, nhưng cách này sẽ làm ứng dụng chậm chạp.

## 2. Tạo kênh khác để bên nhận báo cho bên truyền là "tôi đã nhận đủ, cậu đóng kênh được rồi đó"

Trong ví dụ này tôi bổ xung thêm một kênh kiểu bool `done := make(chan bool)`. Khi receiver đã nhận đủ, không còn dữ liệu nào trong `pipe`, thì receiver sẽ truyền tin để hàm có thể thoát.

`<-done` khiến cho hàm sẽ chờ cho đến khi nhận được tin. Việc truyền vào kênh `done` giá trị `true` hay `false` không quan trọng, mà quan trọng truyền vào đó 1 tin.

```go
func DemoChannel2() {
    pipe := make(chan string)
    done := make(chan bool) //kênh báo khi receiver nhận đủ
    go func() {
        for {
            receiver, more := <-pipe //khi không còn dữ liệu, more sẽ false
            fmt.Println(receiver)
            if !more {
                done <- true //không còn dữ liệu trong pipe, đã nhận đủ !
                return //thoát khỏi go routine
            }
        }
    }()

    pipe <- "water 1"
    pipe <- "water 2"
    pipe <- "water 3"
    pipe <- "water 4"
    close(pipe)
    <-done  //Chờ khi nhận được tin báo từ receiver thì thoát
}
```

### 2.1 Hàm thoát ngay khi nhận được tin `<-done`

Để chứng minh cho việc hàm thoát ngay khi nhận được tin từ kênh `done`, tôi thêm một lệnh trễ và một lệnh in ra terminal ngay sau lệnh `done <- false`. Kết quả lệnh in này sẽ không kịp thực hiện vì hàm đã thoát.

```go
go func() {
    for {
        receiver, more := <-pipe
        fmt.Println(receiver)
        if !more {
            done <- false
            time.Sleep(time.Millisecond)
            fmt.Println("không in ra vì hàm đã thoát")
            return
        }
    }
}()
```

## 3. Tóm lại là

Lập trình đồng thời (concurrent programming) khó hơn lập trình tuần tự rất nhiều. Nếu không sử dụng cơ chế truyền tin bằng channel, chúng ta không thể đảm bảo đồng bộ thứ tự thực thi ở các go routine. Việc dùng hàm trễ `time.Sleep` là một cách hết sức lang băm đi ngược lại tinh thần của Go.

Đi cùng với lập trình đồng thời là vấn đề data racing: tranh chấp đọc /ghi dữ liệu giữa nhiều thread hoặc go routine, dẫn đến dữ liệu không còn toàn vẹn. Thuốc chữa là channel, sychronization bằng Mutex Lock...

Khi gỡ rối ứng dụng có nhiều tác vụ chạy đồng thời, bản thân việc dừng lại để kiểm tra đã phá vỡ tính chất đồng thời của ứng dụng. Do đó bạn nên hiểu rõ bản chất của go routine, channel. Hoặc vì bạn có quá ít thời gian để hiểu thì cứ tiếp tục đọc những bài viết của tôi về Golang và dùng lại code mẫu. Hãy like, share thật mạnh để tôi có động lực code và viết tiếp nhé. THANKS !