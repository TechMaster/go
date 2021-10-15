## Khởi tạo docker container với docker compose

```docker
version: "3.3"

volumes:
  data_db:

services:
  mysql:
    image: mysql:latest
    volumes:
      - data_db:/var/lib/mysql
    ports:
      - 3306:3306
    environment:
      MYSQL_PASSWORD: 123
      MYSQL_ROOT_PASSWORD: 123
      MYSQL_DATABASE: db
```

Chạy file docker-compose

```
docker-compose up -d
```

Hoặc sử dụng docker để khởi tạo container mysql

```docker
docker run -d --name db_mysql -p 3306:3306 \
-e MYSQL_PASSWORD=123 \
-e MYSQL_ROOT_PASSWORD=123 \
-e MYSQL_DATABASE=db \
mysql:latest
```

### Cấu hình trong DBeaver

![](./images/config-dbeaver.png)

Thông số cấu hình điền lấy trong file docker-compose hoặc câu lệnh khởi tạo container và điền vào các ô được khoanh đỏ

Trường hợp gặp lỗi tham khảo link sau: https://community.atlassian.com/t5/Confluence-questions/MySQL-Public-Key-Retrieval-is-not-allowed/qaq-p/778956


## GORM

### Table name

GORM sẽ tự động chuyển tên struct thành tên bảng theo dạng là snake_case

Ví dụ : **Student** -> **students**, **User** -> **users**, **UserDetail** -> **user_details**

```go
type Student struct {
  Id       string `gorm:"primary_key"` // Id là trường mặc định được sử dụng làm khóa chính
  FullName string
  Email    string
  Phone    string
  CardId   string
}
```

Chúng ta có thể thay đổi tên bảng default bằng cách implement interface **Tabler**

```go
type Tabler interface {
  TableName() string
}
```

```go
func(s *Student) TableName() string{
  return "student"
}
```


### Hook (Object Life Cycle)

Hook là các hàm được gọi trước hoặc sau khi thực hiện các thao tác creation, querying, updating, deletion

Chúng ta có thể chỉ định hook cho một model (struct) cụ thể nào đó, nó sẽ tự động được gọi khi creation, querying, updating, deletion

Nếu khi hook được gọi mà xảy ra lỗi, lúc này GORM sẽ rollback transaction

#### Hook tạo object

Khi tạo object chúng ta có 4 Hook sau

```
BeforeSave
BeforeCreate
AfterCreate
AfterSave
```

**Demo Hook BeforeCreate**

Trước khi tạo đối tượng, nếu đối tượng chưa có ID thì tự động tạo ID

```go
func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
  if s.Id == "" {
    fmt.Println("Bổ sung Id cho student")
    s.Id = util.NewID()
  }
  return
}
```

Tương tự chúng có các Hook khi thực hiện **update**, **delete**, **find**

**Hook Update**

```
BeforeSave
BeforeUpdate
AfterUpdate
AfterSave
```

**Hook Delete**

```
BeforeDelete
AfterDelete
```

**Ví dụ**: Nếu student có **id = "1"** thì không được phép xóa

```go
func (s *Student) BeforeDelete(tx *gorm.DB) (err error) {
	if s.Id == "1" {
		return errors.New("user not allowed to delete")
	}
	return
}
```

**Hook Query**

```
AfterFind
```

**Ví dụ**: In ra thông tin của student mỗi khi tìm kiếm

```go
func(s *Student) ToString() {
  fmt.Println("Id : ", s.Id)
  fmt.Println("FullName : ", s.FullName)
  fmt.Println("Email : ", s.Email)
  fmt.Println("Phone : ", s.Phone)
  fmt.Println("CardId : ", s.CardId)
}

func (s *Student) AfterFind(tx *gorm.DB) (err error) {
  s.ToString()
  return
}
```

## Demo query

1. [Tạo bản ghi](./repo/create.go)
2. [Tìm kiếm bản ghi](./repo/select.go)
3. [Update bản ghi](./repo/update.go)
4. [Xóa bản ghi](./repo/delete.go)
5. [Truy vấn với raw query](./repo/raw_query.go)







