Các ngôn ngữ lập trình hầu hết cung cấp module, package để giúp chúng ta làm quen với http server và thực hành với nó

Trước khi bắt đầu, chúng ta cài package realize, package này giúp chúng ta không cần phải chạy lại server mỗi lần thay đổi code. Chi tiết xem [ở đây](https://techmaster.vn/posts/36634/su-dung-package-realize-trong-golang)

### Create simple server

```go
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page"))
}

func DemoServerDefault() {
	http.HandleFunc("/", homeHandle)

	fmt.Println("Server listenning on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000", nil))
}
```

```ListenAndServe(":3000", nil)``` truyền vào tham số thứ 2 là ```nil```, lúc này server sử sử dụng ```DefaultServeMux``` làm Handler chính để xử lý request

**Hỏi:** Vậy có cách nào để thay đổi Handler của server không?

**Trả lời:** : Chúng ta có thể sử dụng ServeMux hoặc custom một Handler

**Ví dụ sử dụng ServeMux**

```go
func aboutPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func DemoServerServeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", aboutPage)

	fmt.Println("Server listenning on port 3000 ...")
	fmt.Println(http.ListenAndServe(":3000", mux))
}
```

ServeMux có phạm vi sử dụng trong local, trong khi DefaultServeMux có phạm vi sử dụng là global, điều này có thể dẫn đến 1 số rủi ro về bảo mật

**Ví dụ sử dụng Custom Handler**

```go
type handle struct{}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Home page"))
	} else if r.URL.Path == "/about" {
		w.Write([]byte("About page"))
	} else {
		w.Write([]byte("Page not found"))
	}
}

func DemoCustomHandle() {
	handle := new(handle)
	fmt.Println(http.ListenAndServe(":3000", handle))
}
```

Chúng ta cần kiểm tra URL để viết các xử lý tương ứng

### Tạo mutilple server

Cách thông thường chúng ta nghĩ đến khi tạo mutilple server như sau

```go
func DemoMutipleServerDefault() {
	fmt.Println("start")
	server1 := createServer(3000)
	fmt.Println(server1.ListenAndServe())
	fmt.Println("done")

	server2 := createServer(3001)
	fmt.Println(server2.ListenAndServe())

	server3 := createServer(3002)
	fmt.Println(server3.ListenAndServe())
}
```

Khi chạy thì chỉ có server ở cổng ```3000``` chạy, trong khi server ở cổng ```3001``` và ```3002``` không hoạt động vì sao lại thế?

Khi chúng ta tạo server 1 và lắng nghe request bằng câu lệnh ```server1.ListenAndServe()``` nó sẽ ```block``` routine hiện tại, trường hợp chúng ta chạy ở func main thì nó sẽ ```block``` routine main, dẫn đến tình trạng những câu lệnh tiếp theo sau câu lệnh này sẽ không được chạy

```Cách giải quyết```: Sử dụng go routine để khởi tạo server

```go
wg := new(sync.WaitGroup)

wg.Add(3)

go func() {
    server := createServer(3000)
    fmt.Println(server.ListenAndServe())
    wg.Done()
}()

go func() {
    server := createServer(3001)
    fmt.Println(server.ListenAndServe())
    wg.Done()
}()

go func() {
    server := createServer(3002)
    fmt.Println(server.ListenAndServe())
    wg.Done()
}()

wg.Wait()
```

Với cách này chúng ta khởi tạo server trong go routine của riêng nó

### Dữ liệu trả về

Chúng ta có thể trả về cho client nhiều kiểu dữ liệu khác nhau:

- Text
- Json
- Xml
- Html
- File
- Template
- ...

Vì vậy chúng ta cần định dạng kiểu dữ liệu trả về trong header của response với thuộc tính **Content-type** trước khi trả về dữ liệu cho client

Ví dụ:

- Text : Content-type : text/plain
- Json : Content-type : application/json
- Xml  : Content-type : application/xml
- Html : Content-type : text/html

### Test API như thế nào
Chúng ta có một số cách để test API:

- Sử dụng trình duyệt
- Postman
- Rest client extension (sử dụng trong VSCode)
- Sử dụng Javascript để tạo request phía client
- Sử dụng Golang tạo request phía backend
- ...

Ví dụ về các viết testing sử dụng Golang để gửi request

```go
func TestHealthCheckHandler(t *testing.T) {
    // Gửi get request đến một url nào đó
	resp, err := http.Get("http://localhost:3000/health")
	if err != nil {
		log.Fatalln(err)
	}

	// Đọc body từ response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert body sang kiểu string
	sb := string(body)
	expected := "OK"

	if sb != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			sb, expected)
	}
}
```

### Query string

Query String là tập hợp các dữ liệu ở dạng key=value mà được thêm vào trong URL đằng sau dấu ```?```, trường hợp có nhiều cặp key=value thì ta sử dụng dấu ```&``` để ngăn cách giữa chúng. Thông thường ta sử dụng Query String để truyền tải dữ liệu lên server, lọc dữ liệu và HTTP method được sử dụng với query string là GET.

Ví dụ về query string

```
http://localhost:3000?name=hung&status=true
http://localhost:3000/users?page=1&limit=10
```

```go
func userHandle(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name, ok := query["name"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing name"))
		return
	}

	// ["a", "b", "c"] => a,b,c
	message := fmt.Sprintf("Hello %s", strings.Join(name, ","))
	w.Write([]byte(message))
}
```

Chú ý ```r.URL.Query()``` trả về kiểu map```(map[string][]string)``` với key là string và value là []string

### Middleware

Middleware là những đoạn mã trung gian nằm giữa các request và response. Nó nhận các request, thi hành các xử lý của riêng nó trước khi đi vào hàm handle chính

```
Router → Middleware Handler → Application Handler
```

Một số trường hợp hay sử dụng middleware như: log request, authentication, authorization, ...

Trong một tiến trình xử lý chúng ta có thể có không hoặc nhiều middleware

Chúng ta có thể sử dụng middleware cho:

- Toàn bộ router
- Một nhóm các router
- Một router cụ thể nào đó

Trong package net/http của golang không hỗ trợ cú pháp giúp chúng ta xử lý với middleware, vì vậy muốn sử dụng middleware thì phải tự custom

Trong trường hợp chúng ta nên sử dụng HOF (Higher Order Function) :
- HOF là hàm nhận một hàm khác vào làm tham số
- Hoặc trả về kết quả là một hàm

Ví dụ

```go
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func DemoMiddleware() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler))) //Higher-Order function

	log.Fatal(http.ListenAndServe(":3000", mux))
}
```

Với ví dụ trên, chúng ta có 2 middleware là ```middlewareOne``` và ```middlewareTwo``` và một main handle là ```final```.

Khi một request đến lúc này nó được xử lý theo trình tự ```middlewareOne -> middlewareTwo -> final```

Trường hợp API là ```/foo```, khi request được xử lý trong ```middlewareTwo``` nó sẽ trả về kết quả luôn cho client mà không cần đi vào ```final``` để xử lý tiếp
