# Lập trình CSDL

Buổi trước, phân tích thiết kế --> thiết kế bảng và quan hệ.

Buổi hôm nay chúng ta học về go-pg ~ ORM (Object Relation Mapping)
JPA - Hibernate tuổi đời 15 năm.

Các pattern phổ biến:
1. Một - Một (rất ít)
2. Một - Nhiều (50%)
3. Nhiều - Nhiều (30-40%)
4. Recurise
5. Polymorphism
6. Phi quan hệ (array, json)

## Các bước thực hành

### Phân cấp thư mục theo quy ước chuẩn
```
├── ddl   -> Data Definition Language
│   └── foo_bar.sql  -> Tạo bảng foo và bar
├── dml   -> Data Manipulation Language
│   └── foo_bar.sql  -> Chèn dữ liệu hoặc xoá trắng bảng foo và bar
├── model -> Lưu các định nghĩa entity
│   └── foo_bar.go
├── repo -> Lưu các hàm thao tác dữ liệu
│   ├── base.go
│   ├── foo_bar.go
│   └── util.go
├── test -> Viết các hàm kiểm thử
│   └── foo_bar_test.go
```

### Tạo database, schema và bảng

Chạy lệnh khởi động Postgresql 14 container
```
docker run --name postgres14 -e POSTGRES_PASSWORD=123 -p 5432:5432 -d postgres:14.0-alpine
```

Dùng DBeaver kết nối vào CSDL

Tạo schema có tên là `test`

```go
type Foo struct {
	tableName struct{} `pg:"test.foo"` //Giống Annotation trong Java MemberShip --> member_ship
	Id        string   `pg:"id,pk"` 
	Name      string
	Bars      []Bar `pg:"rel:has-many"` //Relation mapping kiểu ORM
}
```
1. Tên bảng trong Postgresql nên viết chữ thường các từ nối nhau bằng snake_case
2. Những thuộc tính nào chúng ta sẽ dùng, xem, sửa ở package khác thì cần viết Hoa để công khai nó

#  Ví dụ quan hệ 1:M (một nhiều)

Chúng ta có 2 bảng `foo` và `bar`. Bảng `bar` có trường `foo_id` là foreign key từ bảng `foo.id`

Quan hệ 1:M có thể thấy ở các ví dụ: 
- Một quyển sách (book) có nhiều trang (pages)
- Một người (person) có nhiều địa chỉ (addresses)
- Một đơn hàng (order) có nhiều dòng mặt hàng (line_items)

## Thực hành
Hãy chạy tuần tự các hàm test trong file [foo_bar_test.go](test/foo_bar_test.go)