# Lesson 03

## 1. Struct
Golang không có class mà chỉ có Struct
- Struct không có tính kế thừa.
- Go tổ hợp các struct (composition) thay vì kế thừa (inheritance)
- Struct Golang không có method nhưng có khái niệm receiver. Có 2 loại receiver là Pointer Receiver và Value Receiver. Receiver bản chất là `func`, được viết bên ngoài struct.
- 

```go
type Product struct {
	id    string
	name  string
	price int
}
```

```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  Address
	Address2  Address
}
type Address struct {
	Country string
	City    string
}
```
### Tính chất tổ hợp (composition)

Có thể khai báo Struct chứa bao nhiêu Struct khác tuỳ ý. Đây là Composition

#### Khả năng thứ 2:
```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address1  Address  //value
	Address2  *Address //pointer
}
```

#### Khả năng thứ 3:
```go
type Person struct {
	FirstName string 
	LastName  string
	Age       int
	Addresses []Address //slice chứa các address
}
```

## 2. Pointer
Hỏi: Pointer trong Go khác gì với C hay C++?

Đáp: Go Pointer khác ở những điểm sau đây

Tham khảo bài viết [Pointers in Go](https://dave.cheney.net/2014/03/17/pointers-in-go)
1. No pointer arithmetic (Không có toán tử trên con trỏ)
	```c
	var p *int
	p++
	```
	Tương tự sẽ không có [kỹ thuật con trỏ mảng](https://www.tutorialspoint.com/cprogramming/c_pointer_to_an_array.htm)
2. Go pointer là type safe
	```go
	b := new(int)
	*b = 10
	c := new(int64)	
	c = b  //Khác kiểu nên không thể gán
	```
	Golang kiểm tra kiểu con trỏ ngay lúc compile time, do đó nó là type safe pointer (ngăn trường hợp ứng dụng bị crash vì gán sai kiểu)

3. `string` trong go là immutable value type (kiểu giá trị không thay đổi được sau khi khởi tạo) chứ không phải pointer type. Khác với C, C++, string là reference type.
  Xem [pointer.go](pointer.go)
	```go
	func modifyString(s string) {
		s = s[len(s)-1:] + s[1:len(s)-1] + s[:1]
		fmt.Println("Inside func s = ", s)
	}

	func DemoPointer() {		
		s := "hello"
		modifyString(s)
		fmt.Println("Outside func s = ", s)
	}
	```

	Kết quả
	```
	Inside func s =  oellh
	Outside func s =  hello
	```
	Thực tế khi truyền `func modifyString(s string)` một chuỗi mới được copy rồi truyền vào hàm.

	Để có thể thay đổi được string khi chuyền vào func chúng ta phải làm như sau
	```go
	func modifyString2(s *string) {
		/*	temp := *s
		*s = temp[len(temp)-1:] + temp[1:len(temp)-1] + temp[:1]*/

		*s = (*s)[len(*s)-1:] + (*s)[1:len(*s)-1] + (*s)[:1]
		fmt.Println("Inside func s = ", s)
	}
	```
	Cách khác là trả về string thay đổi
	```go
	func modifyString3(s string) string {
		return s[len(s)-1:] + s[1:len(s)-1] + s[:1]
	}
	```

	Bạn Bền hỏi trong 3 hàm thay đổi chuỗi này, hàm nào chạy nhanh nhất?

	Trả lời: hãy viết hàm Bencmark ! Kết quả đây
	```
	BenchmarkModifyString1-8        44691324                25.58 ns/op
	BenchmarkModifyString2-8        26154003                41.74 ns/op
	BenchmarkModifyString3-8        45897848                25.67 ns/op
	```

	Hoá ra là truyền con trỏ đến chuỗi vào hàm lại chạy chậm. Hy vọng phiên bản kế tiếp Golang sẽ tối ưu được phần này.

	> Lời khuyên của tôi: đừng thay đổi string bằng cách truyền con trỏ *string vừa chậm vừa khó hiểu về cú pháp. Hãy truyền string dạng immutable value và trả về chuỗi thay đổi.

4. Go pointer không có những khái niệm phức tạp như pointer trỏ đến pointer

### 2.1 Keyword `new` để tạo vùng nhớ để pointer trỏ tới

```go
b := new(int)  //cấp phát vùng nhớ kiểu *int
fmt.Printf("b address %p\n", b) //b address 0xc0000c0008
var c *int     //khai báo biến con trỏ kiểu *int
fmt.Printf("c address %p\n", c) //c address 0x0
```
Ví dụ trên c chưa được khởi tạo vùng nhớ

### 2.2 Toán tử `*` và `&`
Toán tử `*` để lấy giá trị mà pointer trỏ đến
```go
b := new(int)
*b = 10

func incIntPointer(a *int) {
	(*a)++
	fmt.Println("Inside func a = ", *a)
}
```

Toán tử `&` để lấy địa chỉ của một đối tượng

```go
x := 10
y := &x //con trỏ y trỏ đến x
fmt.Println("y = ", y)
fmt.Println("*y = ", *y)
```
Ví dụ thay đổi giá trị của a
```go
func incIntPointer(a *int) {
	(*a)++
	fmt.Println("Inside func a = ", *a)  //Inside func a =  11
}

incIntPointer(&a)
fmt.Println("Outside func a = ", a)    //Outside func a =  11
```
### 2.3 In ra địa chỉ của một đối tượng dùng `%p`

Nếu là đối tượng thực sự, phải dùng toán tử `&` để lấy địa chỉ
```go
fmt.Printf("%p\n", &product)
```

Nếu là con trỏ đến đối tượng
```go
fmt.Printf("%p\n", product)
```

### 2.4 Fluent API
Fluent API là kỹ thuật nối chuỗi các hàm cùng thực thi trên con trỏ đến một đối tượng
Tham khảo: 
- [Lập trình Java phong cách Fluent là gì?](https://techmaster.vn/posts/36423/lap-trinh-java-phong-cach-fluent-la-gi)
- [person.go](person.go)

```go
jack := BuildPerson().WithFirstName("Jack").WithLastName("London").WithAge(12)
```

## 3. Value Receiver vs Pointer Receiver
Value receiver không thay đổi giá trị của struct. Thực tế struct sẽ copy ra một struct mới khi truyền vào value receiver
```go
/* Nâng giá sản phẩm lên
price = price * (1 + percentage/100)
Đây là value receiver function
*/
func (product Product) increasePrice1(percentage int) {
	fmt.Printf("%p\n", &product)
	product.price = product.price * (100 + percentage) / 100
}
```

Pointer receiver có thể thay đổi thuộc tính của struct khi kết thúc phương thức
```go
/* Nâng giá sản phẩm lên
price = price * (1 + percentage/100)
Đây là pointer receiver function
*/
func (product *Product) increasePrice2(percentage int) {
	product.price = product.price * (100 + percentage) / 100
}
```
**Tại sao?**
Pointer receiver làm việc trực tiếp trên đối tượng
Value receiver làm việc với bản copy của đối tượng

Kết quả khi chạy
```go
func main() {
	demoProduct()
}
```

```
0xc00010e030 <-- địa chỉ ban đầu của đối tượng product nike 
0xc00010e060 <-- địa chỉ đối tượng bên trong value receiver 
100          <-- giá sản phẩm không đổi
0xc00010e030 <-- địa chỉ đối tượng bên trong pointer receiver
120          <-- giá sản phẩm đã thay đổi
```

### 3.1 Khi nào dùng Pointer khi nào dùng Value?

https://yourbasic.org/golang/pointer-vs-value-receiver/

#### Dùng Pointer Receiver
1. Khi cần thay đổi thuộc tính trong struct
2. Dùng trong hàm có sync.Mutex và sync.WaitGroup
3. Struct lớn, nhiều trường --> việc copy trở nên tốn kém
   
#### Dùng Value Receiver
1. Không muốn thay đổi thuộc tính trong struct
2. Struct đơn giản
3. Kiểu map, func, chan type. Xem [func_type.go](func_type.go)

### 3.2 So sánh tốc độ Pointer receiver vs Value receiver
```
$ go test -bench=.
BenchmarkPassStructAsValue-8            690402608                1.692 ns/op
BenchmarkPassStructAsPointer-8          26917393                40.60 ns/op
BenchmarkValueReceiver-8                634725636                1.880 ns/op
BenchmarkPointerReceiver-8              284371106                4.224 ns/op
```

Nhìn vào kết quả benchmark bạn sẽ thấy:
- truyền struct as value nhanh hơn truyền struct as pointer.
- value receiver chạy cũng nhanh hơn pointer receiver

Vậy chỉ dùng pointer receiver hay truyền pointer khi phải thay đổi giá trị tham số

### 3.3 Thử nghiệm với trường hợp phức tạp hơn

```
├── pointer
│   └── pointer_repo.go
├── value
│   └── value_repo.go
```
```go
func BenchmarkValueAccountRepo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value.InitData()
	}
}

func BenchmarkPointerAccountRepo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointer.InitData()
	}
}

func BenchmarkValueFindById(b *testing.B) {
	value.InitData()
	for i := 0; i < b.N; i++ {
		value.InitData()
	}
}

func BenchmarkPointerFindById(b *testing.B) {
	pointer.InitData()
	for i := 0; i < b.N; i++ {
		pointer.InitData()
	}
}
```

**Kết quả**
```
BenchmarkPointerReceiver-8              283266402                4.233 ns/op
BenchmarkValueAccountRepo-8                 2233            521005 ns/op
BenchmarkPointerAccountRepo-8               2592            460963 ns/op
BenchmarkValueFindById-8                    2542            528298 ns/op
BenchmarkPointerFindById-8                  2481            453237 ns/op
```

Giờ thì pointer lại nhanh hơn value. Vậy phải làm sao bây giờ?

Trả lời: Hãy viết chạy Benchmark ở những hàm quan trọng để chọn phương án tối ưu nhất !

### 3.3 Nên truyền slice dạng value hay pointer
Nếu chúng ta thay đổi giá trị phần tử trong slice bên trong một hàm, khi ra khỏi hàm, thay đổi này vẫn giữ nguyên. Như vậy, truyền slice vào một hàm bản chất là truyền con trỏ rồi đó. Điều này đúng với bản chất của slice là con trỏ đến cấu trúc array.

Có thể dùng con trỏ slice nhưng việc này không cần thiết và có thể còn chậm hơn việc truyền slice thông thường.

## 4. func type ~ Kiểu hàm

Xem [func_type.go](func_type.go)

Trong Golang, chúng ta có thể định nghĩa kiểu hàm `func`, truyền hàm, gán hàm như biến.

```go
type Operator func(a int, b int) int

var op Operator
op = func(a int, b int) int {
	return a + b
}

fmt.Println(op.compute(1, 2))
fmt.Println(op.compute2(1, 2))

op = Subtract
fmt.Println(op.compute(10, 5))
fmt.Println(op.compute2(10, 5))
```

Hỏi: Có dùng được con trỏ hàm không?

Đáp: Được ! dùng khi cần thay đổi nội dung hàm thực thi

```go
type Operator func(a int, b int) int

type POperator *Operator
op = Subtract

var op2 POperator
op2 = &op
fmt.Println((*op2).compute(10, 5))
```

## 5. Escape to Heap

Trong ứng dụng  có 2 loại bộ nhớ: heap và stack. Khi lập trình C, C++, thì biến con trỏ sẽ được cấp phát ở vùng nhớ heap.

Trong Go không phải lệnh `new` nào cũng sẽ được cấp phát vùng nhớ ở heap.

[Golang Escape to Heap - Ví dụ minh hoạ cách Golang cấp phát vùng nhớ](https://techmaster.vn/posts/36541/golang-escape-to-heap-vi-du-minh-hoa-cach-golang-cap-phat-vung-nho)