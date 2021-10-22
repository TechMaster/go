# Xử lý lỗi trong Golang

Các điểm yếu:
1. Không sử dụng try cach exception như Java, C++, C#
2. Không hiển thị được strack trace lỗi ở dòng nào. Rất khó debug
3. Trả về lỗi bằng `retun` nên cấu trúc kiểm tra lỗi khá dài dòng

Điểm cộng:
1. Có tính năng `defer` để thực hiện các thao tác ngay cả khi có lỗi phát sinh
2. Tốc độ xử lý và hỗ trợ tính đa nhiệm có thể chính là nguyên nhân các nhà thiết kế Go sử dụng phương pháp return error 
```go
func Bar() error {
	defer func() {  //hàm nay luôn chạy sau cùng trước bất kỳ lệnh return nào
		fmt.Println("Clean up resource at the end of Bar")
	}()
	return errors.New("Error from Bar function")
}
```
Pacnic hiện thị được stack nhưng hiển thị xong thì chương trình sập luôn.
![](panic.jpg)