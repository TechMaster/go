## Tạo dự án và cấu trúc thư mục - package

```
.
├── README.md
├── controller   -> các hàm xử lý với APi
├── ddl          -> chứa các câu lệnh DDL để tạo bảng
├── go.mod
├── go.sum
├── main.go      -> File chạy chính
├── model        -> định nghĩa model - entity
├── repo         -> Chứa hàm thao tác với cơ sở dữ liệu
├── router       -> Định nghĩa các router
├── test         -> Viết các hàm unit test để kiểm tra repo
├── test_post.rest -> Kiểm thử API post
├── test_user.rest -> Kiểm thử API user
└── tmp
```
## Repo

Đối với mỗi resource chúng ta thực hiện các công việc sau

**User**

- Lấy danh sách user
- Chi tiết user theo ID
- Tạo user
- Cập nhật thông tin user
- Xóa user

**Post**

- Lấy danh sách post ứng với userID
- Chi tiết post theo ID
- Tạo post
- Cập nhật thông tin post
- Xóa post

### Thực hiện Unit Test

Chú ý, khi thực hiện **run debug** unit test với firber-postgres hay iris-postgres cần phải gọi func **ConnectDatabase**

Chúng ta có thể thực hiện như sau

```go
// Cách 1: Gọi function ngay dưới phần định nghĩa
var _ = ConnectDatabase()

// Cách 2: Khai báo func init để thực hiện gọi bên trong
func init() {
	var _ = ConnectDatabase()
}
```

> Trường hợp nếu chạy server fiber/iris thì hãy xóa phần gọi func trên đi, vì đã gọi trong func main rồi

### Thực hiện test API

Công cụ test : **Extension REST Client**

Ví dụ về test API user (file **test_user.rest**)

```go
// Lấy danh sách user
GET http://localhost:3000/users HTTP/1.1

###

// Chi tiết user theo ID
GET http://localhost:3000/users/JpwUSxHf HTTP/1.1


###

// Tạo user mới
POST http://localhost:3000/users HTTP/1.1
Content-Type: application/json

{
    "full_name" : "Ngo Van B",
    "email" : "b@gmail.com",
    "phone" : "0123456789"
}

###

// Cập nhật thông tin user
PUT http://localhost:3000/users/HPD0-A2N HTTP/1.1
Content-Type: application/json

{
    "full_name" : "Tran Thi C",
    "email" : "c@gmail.com",
    "phone" : ""
}

###

// Xóa user
DELETE http://localhost:3000/users/HPD0-A2N HTTP/1.1
```

### Thực hiện test benchmark với fiber-postgres và fiber-mysql

| Trường hợp benchmark | Fiber-Postgres | Fiber-Mysql |
| -------- | -------- | -------- |
| GetAllUser     | 1449, 1446, 1446     | 1104, 1093, 1077     |
| GetUserById     | 1497, 1502, 1508     | 541, 574, 589     |
| CreateUser     | 1452, 1467, 1459     | 206, 219, 217     |
| UpdateUser     | 1396, 1418, 1407     | 308, 288, 306     |