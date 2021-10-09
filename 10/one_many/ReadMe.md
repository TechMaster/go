#  Ví dụ quan hệ 1:M (một nhiều)

Chúng ta có 2 bảng `foo` và `bar`. Bảng `bar` có trường `foo_id` là foreign key từ bảng `foo.id`

Quan hệ 1:M có thể thấy ở các ví dụ: 
- Một quyển sách (book) có nhiều trang (pages)
- Một người (person) có nhiều địa chỉ (addresses)
- Một đơn hàng (order) có nhiều dòng mặt hàng (line_items)

## Thực hành
Hãy chạy tuần tự các hàm test trong file [foo_bar_test.go](test/foo_bar_test.go)