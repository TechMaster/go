# Bài tập buổi số 15

Trong [buổi 15](../../15/ReadMe.md) các bạn đã học:
1. Thiết kế một thư viện [logger](../../15/logger/base.go)
2. Sử dụng logger trong một dự án web phục vụ cả REST và Server Side Rendering [userlogger](../../15/uselogger/ReadMe.md)

Thư viện logger phụ thuộc rất chặt (tightly coupling) vào thư viện iris.
Xem file [log_error.go](../../15/logger/log_error.go)

```go
func Log(ctx iris.Context, err error) {
	if errors.Is(err, syscall.EPIPE) {
		return
	}
	//Trả về JSON error khi client gọi lên bằng AJAX hoặc request.ContentType dạng application/json
	shouldReturnJSON := ctx.IsAjax() || ctx.GetContentTypeRequested() == "application/json"
	switch e := err.(type) {
	case *eris.Error:
		if e.ErrType > eris.WARNING { //Chỉ log ra console hoặc file
			logErisError(e)
		}

		if shouldReturnJSON { //Có trả về báo lỗi dạng JSON cho REST API request không?
			errorBody := iris.Map{
				"error": e.Error(),
			}
			if e.Data != nil { //không có dữ liệu đi kèm thì chỉ cần in thông báo lỗi
				errorBody["data"] = e.Data
			}
			if e.Code > 300 {
				ctx.StatusCode(e.Code)
			} else {
				ctx.StatusCode(iris.StatusInternalServerError)
			}

			_, _ = ctx.JSON(errorBody) //Trả về cho client gọi REST API
			return                     //Xuất ra JSON rồi thì không hiển thị Error Page nữa
		}

		// Nếu request không phải là REST request (AJAX request) thì render error page
		ctx.ViewData("ErrorMsg", e.Error())
		if e.Data != nil {
			if bytes, err := json.Marshal(e.Data); err == nil {
				ctx.ViewData("Data", string(bytes))
			}
		}
		_ = ctx.View(LogConf.ErrorTemplate)
		return
	default: //Lỗi thông thường
		fmt.Println(err.Error()) //In ra console
		if shouldReturnJSON {    //Trả về JSON
			ctx.StatusCode(iris.StatusInternalServerError)
			_, _ = ctx.JSON(err.Error())
		} else {
			_ = ctx.View(LogConf.ErrorTemplate, iris.Map{
				"ErrorMsg": err.Error(),
			})
		}
		return
	}
}
```

## Bài tập số 1
Hãy porting thư viện này logger để nó hoạt động tốt với [fiber](https://github.com/gofiber/fiber)
Chú ý fiber cũng có kiểu context như iris.

## Bài tập số 2
Hãy tìm cách làm sao viết một thư viện Logger có thể dùng cho nhiều web framework khác nhau như Iris, Fiber, Gin.

Định nghĩa ra một interface chung cho các web framework này. Interface gồm một số mẫu hàm căn bản:
1. Kiểm tra request có phải là AJAX , application/json request hay không?
2. Trả về dữ liệu dạng JSON
3. Trả về error page
4. Trả về info page

Nếu bạn nào còn nhớ buổi Design Pattern do anh Bùi Hiên dạy sẽ làm được.
[https://refactoring.guru/design-patterns/go](https://refactoring.guru/design-patterns/go)
