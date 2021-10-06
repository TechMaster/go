package repo

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"gopgexample/goccy"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/pgjson"
)

var (
	DB     *pg.DB     //kết nối vào CSDL Postgresql
	random *rand.Rand // Đối tượng dùng để tạo random number
)

/*
Hàm init này luôn chạy đầu tiên của package repo
*/
func init() {
	//Mở kết nối vào CSDL Postgresql
	DB = pg.Connect(&pg.Options{
		Addr:     "techmaster3:5432",
		User:     "code123",
		Password: "123456",
		Database: "demo",
	})

	//Thay thế thư viện encoding/json chuẩn bằng thư viện goccy json chạy nhanh hơn
	pgjson.SetProvider(goccy.NewJSONProvider()) //Sử dụng goccy json

	//Log các câu lệnh SQL thực thi để debug
	DB.AddQueryHook(dbLogger{}) //Log query to console

	//Khởi động engine sinh số ngẫu nhiên
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

	//Đăng ký bảng quan hệ nhiều - nhiều
	//orm.RegisterTable((*model.MemberClub)(nil))
}

type dbLogger struct{}

// Hàm hook (móc câu vào lệnh truy vấn) để in ra câu lệnh SQL query
func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// Hàm hook chạy sau khi query được thực thi
func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	bytes, _ := q.FormattedQuery()
	fmt.Println(string(bytes))
	return nil
}
