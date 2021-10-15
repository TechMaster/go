package repo

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	host     string = "localhost"
	port     string = "3306"
	username string = "root"
	password string = "123"
	database string = "db"
)

var (
	DB     *gorm.DB
	random *rand.Rand // Đối tượng dùng để tạo random numbe
)

func ConnectDB() *gorm.DB{
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect database")
	}

	//Khởi động engine sinh số ngẫu nhiên
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)

	return DB
}
