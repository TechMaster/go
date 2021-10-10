# Quan hệ User - Role

Sử dụng cột mảng hay dùng quan hệ nhiều nhiều sẽ tốt hơn?

##  Hướng dẫn chạy thử ví dụ
###  1. Chỉnh sửa kết nối đến cơ sở dữ liệu
Vào file [repo/base.go](repo/base.go) để chỉnh lại cấu hình kết nối CSDL

```go
func init() {
	//Mở kết nối vào CSDL Postgresql
	DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "123",
		Database: "gopg",
	})

	//Log các câu lệnh SQL thực thi để debug
	DB.AddQueryHook(dbLogger{}) //Log query to console

	//Khởi động engine sinh số ngẫu nhiên
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

}
```
Thay đổi pg.Options bằng cấu hình database của bạn

### 2. Tạo bảng bằng DDL scripts
Chạy lệnh tạo bảng dùng công cụ DBeaver hoặc PgAdmin bằng các DDL scripts trong file [ddl/user_role.sql](ddl/user_role.sql)
Tạo kiểu enum:
```sql
CREATE TYPE test.rolenum AS ENUM (
'ADMIN',
'STUDENT',
'TRAINER',
'SALE',
'EMPLOYER',
'AUTHOR',
'EDITOR',
'MAINTAINER');
```
Tạo bảng user:
```sql
DROP TABLE IF EXISTS test.users;
CREATE TABLE test.users (
	id text PRIMARY KEY,
	name text NOT NULL,
	email text,
	mobile text,
	int_roles int[],
	enum_roles test.rolenum[]
);
```
Tạo bảng roles:
```sql
DROP TABLE IF EXISTS test.roles;
CREATE TABLE test.roles (
	id int PRIMARY KEY,
	name text NOT NULL UNIQUE
);
```
Chèn dữ liệu vào bảng roles:
```sql
INSERT INTO test.roles (id, name) VALUES (1, 'ADMIN');
INSERT INTO test.roles (id, name) VALUES (2, 'STUDENT');
INSERT INTO test.roles (id, name) VALUES (3, 'TRAINER');
INSERT INTO test.roles (id, name) VALUES (4, 'SALE');
INSERT INTO test.roles (id, name) VALUES (5, 'EMPLOYER');
INSERT INTO test.roles (id, name) VALUES (6, 'AUTHOR');
INSERT INTO test.roles (id, name) VALUES (7, 'EDITOR');
INSERT INTO test.roles (id, name) VALUES (8, 'MAINTAINER');
```
Tạo bảng user_roles:
```sql
DROP TABLE IF EXISTS test.user_role;
CREATE TABLE test.user_role (
   user_id text REFERENCES test.users(id),
   role_id int REFERENCES test.roles(id)
);
```
### 3. Định nghĩa go-pg model tương ứng với bảng
Trong file [model/user_role](model/user_role) chúng ta sẽ định nghĩa struct `User` là model tương ứng với bảng user trong database
```go
type User struct { //Tên Entity --> biên thành snake_case trong Postgresql
	tableName  struct{} `pg:"test.users"` //--> trỏ đến schema.table
	Id         string   `pg:"id,pk"`      //pk--> đây là primary key
	Name       string   //-> name, kiểu string --> text
	Email      string
	Mobile     string
	Int_roles  []int    `pg:"int_roles,array"`  //Quy ước IntRoles --> int_roles snake case
	Enum_roles []string `pg:"enum_role,array"` //kiểu cột là array lưu string
}
```
Tương tự với model UserRoles:
```go
type User_Role struct {
	tableName struct{} `pg:"test.user_role"`
	User_id   string
	Role_id   int
}
```

### 4. Sử dụng Go-pg để thao tác dữ liệu với database
Chúng ta sẽ viết những đoạn DML bằng Go-pg trong thư mục [repo](repo)
Ví dụ tạo fake user bằng hàm `Create_fake_user()` :
```go
func Create_fake_user(trans ...*pg.Tx) (user *model.User, err error) {
	var transaction *pg.Tx
	// kiểm tra xem có tham số trans hay không? Nếu có tham số trans thì dùng tham số trans làm transaction, còn không tự tạo mới transaction
	if len(trans) > 0 {
		transaction = trans[0]
	} else {
		transaction, err = DB.Begin()
		if err != nil {
			return nil, err
		}
	}

	// Tạo random Id và Roles
	id, _ := gonanoid.New(8)
	int_roles, enum_roles := gen_random_roles(3)

	// Gán các giá trị random vào model
	user = &model.User{
		Id:         id,
		Name:       gofakeit.Name(),
		Email:      gofakeit.Email(),
		Mobile:     gofakeit.Phone(),
		Int_roles:  int_roles,
		Enum_roles: enum_roles,
	}

	// Chèn thêm data qua model user vào bảng user bằng go-pg
	_, err = transaction.Model(user).Insert()
	if !check_err(err, transaction) {
		return nil, err
	}

	// Lấy thông tin role và id của user để chèn data vào bảng user_roles tương tự với bảng user
	for _, role := range int_roles {
		user_role := model.User_Role{
			User_id: id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !check_err(err, transaction) {
			return nil, err
		}
	}

	//Nếu dùng transaction tạo trong hàm này thì commit luôn
	if len(trans) == 0 {
		if err = transaction.Commit(); err != nil {
			return nil, err
		} else {
			return user, nil
		}
	} else {
		return user, nil
	}
}
```

### 5. Viết unit test kiểm tra kết quả
Chúng ta sẽ viết các đoạn hàm unit test trong thư mục [test](test)
Ví dụ test hàm `Create_fake_user()` trong repo:
```go
func Test_create_fake_user(t *testing.T) {
	user, err := repo.Create_fake_user()
	assert.Nil(t, err)
	if err == nil {
		fmt.Println(user.Id, user.Name)
	}
}
```
run test để kiểm tra kết quả.
## Bài tập
1. Hãy hoàn thành các hàm trong file [repo/user_role_query.go](repo/user_role_query.go)
2. Bổ xung hàm bổ xung, xoá roles trong file [repo/user_role_update.go](repo/user_role_update.go)