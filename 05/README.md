## Design Pattern

1. **Design Pattern là gì**

- Tập hợp các giải pháp giúp chúng ta có thể giải quyết các vấn đề thường xuyên gặp phải khi thiết kế phần mềm
- Không phải là các đoạn code có sẵn mà là mô tả cách giải quyết một vấn đề mà có thể được dùng trong nhiều tình huống khác nhau
- Áp dụng cho hầu hết các ngôn ngữ hiện này: Java, Python, C, C++, Go, ...
- Design pattern có thể giúp thiết kế của chúng ta linh hoạt, dễ dàng thay đổi và bảo trì hơn.

2. **Các nhóm Design Pattern**

- Nhóm khởi tạo (Creational Pattern): bao gồm các giải pháp liên quan đến khởi tạo đối tượng
    - Singleton
    - Factory
    - Factory method
    - Builder
    - Prototype

- Nhóm cấu trúc (Structural Pattern): thiết lập liên kết, mối quan hệ giữa các đối tượng
    - Adapter
    - Bridge
    - Composite
    - Proxy
    - ...

- Nhóm hành vi (Behavioral Pattern): Liên quan đến hành vi của đối tượng và cách các đối tượng giao tiếp với nhau
    - State
    - Visitor
    - Command
    - Observer
    - ...

3. **Tại sao cần sử dụng Design Pattern**

- Giúp dự án của chúng ta dễ bảo trì, nâng cấp, mở rộng
- Hạn chế rủi ro tiềm ẩn
- Dễ dàng làm việc với các thành viên trong team

4. **Cách học design pattern?**

- Hiểu mục đích của pattern này dùng để làm gì? Áp dụng khi nào?
- Diagram UML
- Code ví dụ triển khai

## Anti Design Pattern

Đây cũng là những pattern giúp chúng ta giải quyết vấn đề nhưng qua thời gian thì chúng trở lên có hại nhiều hơn lợi

Một số ví dụ điển hình của Anti Design Pattern

- Hard code
- Magic number and string
- Copy and paste programming
- Dependency Hell
- Spaghetti Code
- Tái phát minh
- ..

## Singleton pattern

Singleton Pattern đảm bảo rằng chỉ có duy nhất một instance, và cung cấp cách thức để chúng ta có thể truy cập tới instance bất cứ lúc nào.

Một số ví dụ sử dụng singleton pattern:

- Database
- Logger
- Config
- File IO

### Một số cách triển khai

- Hàm khởi tạo : Chú ý đến trường hợp đa luồng
- Sử dụng sync.Once (recommend)
- Sử dụng Mutex Lock
- Sử dụng hàm init của package : Đối tượng sẽ bắt buộc phải khởi tạo khi khởi tạo package

Kết nối đến database

```go
func DemoDatabase() {
	db := GetInstance()

    // Đóng kết nối đến database khi hàm kết thúc
	defer db.Close()

	// Kiểm tra kết nối
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	// Tìm kiếm thông tin user
	var user User
	err := db.Model(&user).Where("id=?", "1234").Select()
	// SELECT * FROM users WHERE id = '1234'

	if err != nil {
		fmt.Println("Read Error")
	}
}
```

> Best Practice : Hãy sử dụng defer ngay sau khi tạo kết nối đến đối tượng (database, file, ...)

**Singleton + Interface : Khá là quan trọng**

Ngăn không cho truy xuất trực tiếp vào các thuộc tính của object nhằm mục đích thay đổi

**Ví dụ:** Một tài khoản ngân hàng mà cho người dùng truy xuất trực tiếp để thay đổi thông tin rất nguy hiểm

Hãy tạo interface và định nghĩa các hành vi mà người dùng có thể thao tác với tài khoản đó

```go
type BankAccount interface {
    ShowInfo()
    PayBills()
    TransferMoney()
}

type Account struct {
    Name string
    Email string
    Money int
}

func(a Account) ShowInfo() { /* code xử lý */ }

func(a Account) PayBills() { /* code xử lý */ }

func(a Account) TransferMoney() { /* code xử lý */ }
```

## Prototype pattern

### Định nghĩa

Prototype pattern sẽ tạo ra một object mới bằng cách copy, object mới sẽ mang đầy đủ đặc tính của object ban đầu, và chúng ta có thể thay đổi object mới mà không ảnh hưởng gì đến object ban đầu

### Ví dụ thực tế

Bây giờ chúng ta có 1 đoạn văn dài khoảng 1 trang

Nhiệm vụ của chúng ta là tạo ra 1 đoạn văn khác y hệt đoạn văn đầu

Chúng ta không thể nào ngồi gõ lại đoạn văn được (1 đoạn văn thì được nhưng 100 đoạn văn chắc bó tay rồi)

⇒ Giải pháp của chúng ta là: Copy and paste

Đoạn văn thứ 2 muốn viết thêm thì cũng không ảnh hưởng gì đến thuộc tính ban đầu

### Các thành phần cơ bản

Một Prototype Pattern gồm các thành phần cơ bản sau:

- **Prototype** : khai báo interface cho việc clone chính nó.
- **ConcretePrototype** : phần này sẽ implement cho method được khai báo ở interface
- **Client** : có thể copy bất kì object nào thông qua phương thức clone

### Sử dụng khi nào

- Được dùng khi việc tạo một object tốn nhiều chi phí và thời gian trong khi bạn đã có một object tương tự tồn tại.
- Muốn truyền đối tượng vào một hàm nào đó để xử lý, thay vì truyền đối tượng gốc có thể ảnh hưởng dữ liệu thì ta có thể truyền đối tượng sao chép.
- Ẩn độ phức tạp của việc khởi tạo đối tượng từ phía Client (phía người sử dụng).

### **Ưu và nhược điểm của pattern này**

**Ưu điểm**

- Có thể dễ dàng sao chép 1 đối tượng mà không cần quan tâm đến struct cụ thể
- Có thể loại bỏ các đoạn code lặp đi lặp lại khi khởi tạo đối tượng
- Có thể tạo ra các đối tượng phức tạp một cách thuận tiện và nhanh chóng

**Nhược điểm**

- Việc sao chép các đối tượng phức tạp có nhiều tham chiếu đối khi rất phức tạp

```jsx
func(p *Person) Clone() iPerson{
	return &Person{
		Name: p.Name,
		Email: p.Email,
		Age: p.Age,
		Address: p.Address.Clone(),
	}
}

func(a Address) Clone() *Address{
	return &Address{
		City: a.City,
		Stress: a.Stress,
	}
}
```

## Builder design pattern

Cho phép xây dựng một đôi tượng phức tạp bằng cách sử dụng các đối tượng đơn giản và sử dụng hình thức tiếp cận từng bước

**Ví dụ**: chế tạo Robot, việc chế tạo sẽ được thực hiện từng bước 1 như: lắp tay, lắp chân, lắp đầu, tô màu, ... -> ra đối tượng hoàn chỉnh

## **Khi nào thì sử dụng Builder Pattern?**

Dưới đây là một số trường hợp bạn nên cân nhắc sử dụng **builder pattern** cho code của mình:

- Các package của bạn chứa quá nhiều hàm khởi tạo hoặc những hàm khởi tạo quá cồng kềnh và phức tạp.
- Bạn không muốn việc gán giá trị cho các tham số của hàm khởi tạo phải tuân theo một trật tự cố định nào đó, ví dụ: Thay vì phải gán giá trị tuần tự từ tham số A rồi mới đến tham số B và tham số C, bạn có thể gán giá trị cho tham số B trước rồi mới đến A và C.

## **Cách thức triển khai thế nào?**

Chúng ta có thể thấy, **Builder Pattern** sẽ gồm có **4 thành phần chính**:

- **Product** : đại diện cho đối tượng cần tạo, đối tượng này phức tạp, có nhiều thuộc tính.
- **Builder** : là interface khai báo phương thức tạo đối tượng.
- **ConcreteBuilder** : implement Builder và định nghĩa chi tiết cách tạo ra đối tượng.
- **Director**: là nơi sẽ gọi tới Builder để tạo ra đối tượng.