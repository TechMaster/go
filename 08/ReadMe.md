# Giới thiệu các web framework khác nhau

Trong bài này chúng ta sẽ làm quen với những web framework nổi tiếng, có rating cao, sử dụng được ngay cho production.

1. [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
2. [https://github.com/gofiber/fiber](https://github.com/gofiber/fiber)
3. [https://github.com/labstack/echo](https://github.com/labstack/echo)
4. [https://github.com/kataras/iris](https://github.com/kataras/iris)

> Đã có code mẫu ở [mygin](mygin) và [myiris](myiris)

### Những tiêu chí đánh giá:

1. Tốc độ trả về plain text 'Hello World'
2. Tốc độ trả về cấu trúc JSON dạng mảng 20 phần tử struct
3. Điều hướng
   - Gom nhóm các URL chung gốc
   - Xử lý tham số
   - Hỗ trợ regex trong  path
4. Sử dụng middleware: logger, recoveryy
5. Parse request JSON
6. Bạn rút ra kết luận gì từ kết quả benchmark

### Sinh viên thực hành
Một nửa lớp porting ví dụ sang fiber và một nửa lớp porting ví dụ sang echo

Tiếp tục benchmark để rút ra kết luận.

Vấn đề đặt ra là khi chúng ta vào một tổ chức. Chúng ta luôn mặc nhiên tuân thủ, dùng lại những framework có sẵn, nhưng ít khi đặt câu hỏi `why`:

1. Hệ thống đã thực sự tối ưu chưa?
2. Có cần thiết phải mua một phần mềm đắt đỏ, mã nguồn đóng khi có nhiều phần mềm mã nguồn mở tốt hơn?
3. Khi chúng ta dùng một thư viện, nếu thư viện đó có lỗi, tác giả không sửa, chúng ta có tự sửa được không?
4. Khi chọn sử dụng một thư viện/phần mềm chúng ta đã viết kiểm thử tất cả các hard case benchmark thư viện này chưa hay chỉ đọc review trên mạng rồi lập trình thử rồi dùng luôn?