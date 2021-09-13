package email

import (
	"fmt"
	"net/smtp"
)

type GmailSTMP struct {
	config *SMTPConfig
}

func InitGmail(config *SMTPConfig) {
	Emailer = GmailSTMP{
		config: config,
	}
}

//--- Hai phương thức implement interface MailSender
func (gmail GmailSTMP) SendPlainEmail(to []string, subject string, body string) error {

	emailAuth := smtp.PlainAuth("me", gmail.config.From, gmail.config.Password, gmail.config.Host)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%d", gmail.config.Host, gmail.config.Port)

	if err := smtp.SendMail(addr, emailAuth, gmail.config.From, to, msg); err != nil {
		return err
	}
	return nil
}

func (gmail GmailSTMP) SendHTMLEmail(to []string, subject string, data map[string]interface{}, layout ...string) error {
	emailAuth := smtp.PlainAuth("me", gmail.config.From, gmail.config.Password, gmail.config.Host)

	body, err := renderHTML(data, layout...)
	if err != nil {
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%d", gmail.config.Host, gmail.config.Port)

	if err := smtp.SendMail(addr, emailAuth, gmail.config.From, to, msg); err != nil {
		return err
	}
	return nil
}
