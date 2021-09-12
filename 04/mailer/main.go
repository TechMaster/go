package main

import (
	"fmt"
	"mailer/email"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-pg/pg/v10"
)

func main() {
	//sendFakeMail()
	sendEmailToDB()
}

func sendFakeMail() {
	email.InitFakeGmail(&email.SMTPConfig{
		Host:     "smtp.gmail.com",
		From:     "testechmaster@gmail.com",
		Password: "hanoijava123-", //Đừng đổi password. Làm ơn !
		Port:     587,
	}, "cuong@techmaster.vn") //Đổi thành địa chỉ email của bạn để nhận được email này

	err := email.Emailer.SendPlainEmail([]string{"john@gmail.com"}, "Hôm nay ăn gì?", "Bún đậu mắm tôm và bia")
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{
		"title":    "Anh",
		"customer": "Trịnh Minh Cường",
		"foods":    []string{"Cafe nâu đá", "Bánh Croissant", "Olive Ý", "Cá hồi hun khói"},
	}
	err = email.Emailer.SendHTMLEmail([]string{"hana@gmail.com"}, "Hôm nay ăn gì?", data, "food.html")
	if err != nil {
		fmt.Println(err)
	}
}

func sendEmailToDB() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "123",
		Database: "postgres",
		Addr:     "localhost:5432",
	})

	email.InitEmailDB(db)
	defer db.Close()

	for i := 0; i < 10; i++ {
		err := email.Emailer.SendPlainEmail([]string{gofakeit.Email()}, gofakeit.Sentence(5), gofakeit.Paragraph(2, 4, 4, "/n/n"))
		if err != nil {
			fmt.Println(err)
		}
	}

	//DELETE FROM emailstore để xoá bớt dữ liệu

}
