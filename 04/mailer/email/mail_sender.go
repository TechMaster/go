package email

import (
	"bytes"

	"html/template"
)

type MailSender interface {
	SendPlainEmail(to []string, subject string, body string) error
	SendHTMLEmail(to []string, subject string, data map[string]interface{}, layout ...string) error
}

type SMTPConfig struct {
	Host     string
	From     string
	Password string
	Port     int
}

var Emailer MailSender //Đối tượng phụ trách gửi email

var defaultEmailLayout string //Layout mặc định cho email

func SetDefaultEmailLayout(defaultLayout string) {
	defaultEmailLayout = defaultLayout
}

/*
Sinh HTML cho email body

*/
func renderHTML(data map[string]interface{}, layout ...string) (string, error) {
	var emailLayout string

	if len(layout) == 0 {
		emailLayout = defaultEmailLayout
	} else {
		emailLayout = layout[0]
	}

	tmpl := template.Must(template.ParseFiles(emailLayout))

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
