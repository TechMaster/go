# Xử lý lỗi trong Golang

## 1. Golang xử lý lỗi như thế nào
1. Golang không có try catch exception, chỉ có hàm trả về lỗi
2. Một khi đã viết hàm Go bạn cần phải quyết định hàm này có trả về lỗi hay không?  90% hàm phải trả về lỗi
3. Lỗi được tạo ra để xử lý và cần phải được xử lý lỗi đến nơi đến trốn. Tuyệt đối không được dập lỗi, lờ đi.
4. Nếu một hàm trả về nhiều tham số, thì tham số lỗi luôn để cuối cùng
	```go
	func Foo() (result string, count int, err error)
	```
4. Bất kỳ lỗi nào trong Golang đều phải tuân thủ interface
	```go
	type error interface {
    	Error() string
	}	
	```

## 2. Xử lý lỗi
### 2.1 Các bước làm việc với lỗi
1. Define function return error: Định nghĩa hàm trả về lỗi
2. Create error: Tạo error phù hợp
3. Handle error: Xử lý lỗi gồm có kiểm tra loại lỗi, mức độ lỗi, mã lỗi
4. Report error: Báo lỗi cho client, cần quyết định nội dung chi tiết đến mức nào và kiểu báo lỗi. Có hai kiểu báo lỗi:
	- Ứng dụng Server Side Rendering thì trả về trang báo lỗi error page
	- Ứng dụng Client Side, Mobile thì trả về JSON error cùng HTTP status code 
5. Log error: Lỗi nghiêm trọng cần được in ra màn hình console và ghi ra file. Với lỗi panic bắt buộc dừng chương trình bằng hàm `panic("error message")`


### 2.2 Căn bản về lỗi
Một lỗi đầy đủ cần có:
1. Mô tả lỗi: "Failed to connect database", "User is not found"
2. Cấp độ lỗi: `WARNING, ERROR, SYSERROR, PANIC` quyết định cách thức dev báo cáo lỗi và log lỗi.
   - `WARNING`: cảnh báo, thường cho client
   - `ERROR`: các lỗi phía server, ví dụ sai logic
   - `SYSERROR`: các lỗi phía server, lỗi hệ thống, có thể khắc phục
   - `PANIC`: lỗi phía server, lỗi hệ thống, không thể khắc phục, cần dừng chương trình
3. Stack Trace danh sách các hàm gọi nhau gây ra lỗi
4. HTTP Status Code nếu là lỗi sẽ trả về cho REST Client
5. Dữ liệu bổ trợ cho lỗi

Những hành động của lập trình với lỗi:
1. Báo cáo lỗi cho client: trả về trang báo lỗi dễ hiểu, thân thiện
2. Trả về lỗi dạng JSON đối với REST API request
3. In lỗi ra màn hình terminal, sẽ bị mất khi docker container nâng cấp
4. Ghi lỗi vào file, bền vững hơn
5. Bỏ qua lỗi nếu thấy cần (hãn hữu thôi nhé)
6. Nâng cấp độ lỗi lên mức cao hơn
7. Tạo ra một lỗi từ một lỗi khác để thêm thông báo, và dữ liệu bổ trợ

Không xử lý lỗi đúng dẫn đến vấn đề gì?
1. Người dùng không hiểu chuyện gì đã xảy ra
2. Lập trình viên không dò vết (không xem được Stack Trace của lỗi), vì lỗi qua chung chung, khó hiểu
3. Hệ thống sập vì lỗi không được xử lý đúng, chương trình chạy tiếp với biến rỗng (nil)

### 2.3 Log lỗi
Cần phân biệt rõ báo lỗi và log lỗi. Báo lỗi dùng để báo cho client, người dùng cuối. Còn log lỗi là cho hệ thống nội bộ và lập trình viên debug, fix lỗi. Do đó Log lỗi phải chi tiết đầy đủ, cần gồm cả stack trace và thông tin hoàn cảnh lỗi phát sinh. Ngược lại báo lỗi cần ưu tiên sự thân thiện với người dùng.

Log lỗi sẽ có 2 cấp độ:
1. In ra màn hình console
2. Ghi vào file log

Bạn có thể sử dụng các hàm thông thường của Golang hay một thư viện như Uber Zap để log lỗi. Tránh viết logic log lỗi ở mọi nơi khiến code vừa dài, mà vừa phụ thuộc chặt (tightly coupling) vào một thư viện báo lỗi bên thứ ba. Nên tận dụng một hàm xử lý lỗi chung gắn với ứng dụng. Muốn được như vậy, ta phải tuần tự trả về lỗi từ hàm con ra hàm cha, từ hàm cha ra hàm ông, cụ, kỵ...