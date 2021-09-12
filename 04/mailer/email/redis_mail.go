package email

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/spf13/viper"

	"github.com/hibiken/asynq"
)

const (
	SEND_EMAIL = "email:send"
)

type EmailPayload struct {
	To      []string
	Subject string
	Msg     []byte
}

type RedisMail struct {
}

var asynqClient *asynq.Client

func InitRedisMail() *asynq.Client {
	asynqClient = asynq.NewClient(asynq.RedisClientOpt{
		Network:  viper.GetString("redis.network"),
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       1, //Do not use 0 because
	})

	Emailer = RedisMail{}
	return asynqClient
}

/*
Implement MailSender interface
*/
func (rmail RedisMail) SendPlainEmail(to []string, subject string, body string) error {

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)

	payload, err := json.Marshal(EmailPayload{
		To:      to,
		Subject: subjectStr,
		Msg:     msg,
	})
	if err != nil {
		return err
	}

	info, err := asynqClient.Enqueue(asynq.NewTask(SEND_EMAIL, payload))
	if err != nil {
		return err
	}
	fmt.Printf("enqueued task: id=%s queue=%s\n", info.ID, info.Queue)
	return nil
}

func (rmail RedisMail) SendHTMLEmail(to []string, subject string, data map[string]interface{}, layout ...string) error {
	body, err := renderHTML(data, layout...)
	if err != nil {
		return err
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectStr := "Subject: " + subject + "!\n"
	msg := []byte(subjectStr + mime + "\n" + body)

	payload, err := json.Marshal(EmailPayload{
		To:      to,
		Subject: subjectStr,
		Msg:     msg,
	})
	if err != nil {
		return err
	}

	info, err := asynqClient.Enqueue(asynq.NewTask(SEND_EMAIL, payload))
	if err != nil {
		return err
	}
	fmt.Printf("enqueued task: id=%s queue=%s\n", info.ID, info.Queue)
	return nil
}
