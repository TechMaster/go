package email

import (
	"time"

	"github.com/go-pg/pg/v10"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type EmailDB struct {
	db *pg.DB
}

func InitEmailDB(db_ *pg.DB) {
	Emailer = EmailDB{
		db: db_,
	}

}

type EmailStore struct {
	tableName  struct{} `pg:"emailstore"`
	Id         string
	Receipient string
	Subject    string
	Body       string
	CreatedAt  time.Time
}

func (emailDB EmailDB) SendPlainEmail(to []string, subject string, body string) error {
	id, _ := gonanoid.New(8)
	emailitem := EmailStore{
		Id:         id,
		Receipient: to[0],
		Subject:    subject,
		Body:       body,
	}
	if _, err := emailDB.db.Model(&emailitem).Insert(); err != nil {
		return err
	}
	return nil
}

func (emailDB EmailDB) SendHTMLEmail(to []string, subject string, data map[string]interface{}, layout ...string) error {
	body, err := renderHTML(data, layout...)
	if err != nil {
		return err
	}
	id, _ := gonanoid.New(8)
	emailitem := EmailStore{
		Id:         id,
		Receipient: to[0],
		Subject:    subject,
		Body:       body,
	}
	if _, err := emailDB.db.Model(&emailitem).Insert(); err != nil {
		return err
	}
	return nil
}

func (emailDB EmailDB) GetMail(to string) (email *EmailStore, err error) {
	email = new(EmailStore)
	if _, err := emailDB.db.Query(email, `SELECT * FROM debug.emailstore 
	WHERE receipient = ?
	ORDER BY id DESC LIMIT 1`, to); err != nil {
		return nil, err
	}
	return email, nil
}
