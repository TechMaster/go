package database

import (
	"context"
	"fmt"
	"singleton-pattern/config"

	"github.com/go-pg/pg/v10"
)

var database *pg.DB

func GetInstance() *pg.DB {
	if database == nil {
		return connectDatabase()
	}
	return database
}

func connectDatabase() *pg.DB {
	conf := config.GetInstance()
	db := pg.Connect(&pg.Options{
		Addr:     conf.Database.Addr,
		User:     conf.Database.User,
		Password: conf.Database.Password,
		Database: conf.Database.Database,
	})

	return db
}

// DEMO

type User struct {
	/* Table name */
	tableName struct{} `pg:"users"`
	Id        string    `pg:",pk"`
	FullName  string
	Email     string
}


func DemoDatabase() {
	db := GetInstance()
	defer db.Close()

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	// Tìm kiếm thông tin user
	var user User
	err := db.Model(&user).Where("id=?", "1234").Select()

	if err != nil {
		fmt.Println("Read Error")
	}

	showDetail(user)

	// Thêm user vào cơ sở dữ liệu
	_, err = db.Model(&User{
        Id:   "3333",
        FullName: "Trần Văn C",
		Email: "c@gmail.com",
    }).Insert()

    if err != nil {
        fmt.Println("Insert Error")
    }
}

func showDetail(user User) {
	fmt.Printf("Name : %s \n", user.FullName)
	fmt.Printf("Email : %s \n", user.Email)
}

