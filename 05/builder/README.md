## Builder Pattern

### Builder Pattern là gì?
Builder pattern là một trong những Design Pattern thuộc nhóm Creational Design Pattern 

Mẫu thiết kế này cho phép chúng ta tạo ra những đối tượng phức tạp thông qua các câu lệnh đơn giản để tác động nên các thuộc tính của nó.

### Khi nào thì sử dụng Builder Pattern?

- Hàm khởi tạo quá cồng kềnh và phức tạp.

- Chúng ta không muốn việc gán giá trị cho các tham số của hàm khởi tạo phải tuân theo một trật tự cố định nào đó

### Cách thức triển khai thế nào?

Hình ảnh minh họa về cách triển khai build design pattern

![](./doc/image/builder-uml.png)

Chúng ta có thể thấy, Builder Pattern sẽ gồm có 4 thành phần chính:

- Product : đại diện cho đối tượng cần tạo, đối tượng này phức tạp, có nhiều thuộc tính.

- Builder : Định nghĩa interface khai báo phương thức tạo đối tượng.

- ConcreteBuilder : implement Builder và cài đặt chi tiết cách tạo ra đối tượng.

- Director: là nơi sẽ gọi tới Builder để tạo ra đối tượng.
