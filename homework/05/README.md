**Bài 1**

Triển khai bài số 2 (Docker), phần bài tập buổi 03 theo mô hình singleton pattern + kết hợp với interface

Ví dụ

```go
type iDockerClient interface {
    ListAll()
    StartContainer()
    StopContainer()
    ...
}
```

Bên trên chỉ là ví dụ về interface, khi triển khai function prototype có thể có tham số truyền vào hoặc kết quả trả về, mọi người tự liệt kê vào

Ví dụ

```go
type iDockerClient interface {
    ...
    StopContainer(string) error
}
```

**Bài 2**

Chúng ta có 2 loại nhân viên trong công ty là Permanent (nhân viên chính thức) và Contract (nhân viên hợp đồng) được định nghĩa bằng kiểu struct.

```go
type Permanent struct {
    empId    int
    basicpay int
    pf       int
}

type Contract struct {
    empId  int
    basicpay int
}
```

Mức lương của nhân viên **Permanent** là tổng của **basicpay** và **pf** còn đối với nhân viên **Contract** thì chỉ là **basicpay**

Yêu cầu:

- Áp dụng Prototype Pattern để thực hiện việc tạo ra 1 danh sách nhân việc thuộc 2 loại trên (số lượng tùy ý), lưu ý **empId** phải khác nhau
- Tính tổng số tiền mà công ty phải trả cho danh sách nhân viên đã tạo ở trên

(Hãy áp dụng interface cho yêu cầu tính lương kể trên)