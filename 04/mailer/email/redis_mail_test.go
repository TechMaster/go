package email

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

/*
Kiểm thử  chức năng  gửi  email qua Redis Stream và AsynQ
github.com/hibiken/asynq
Hãy chạy ở chế độ debug test
*/
func Test_Send_Email_Send(t *testing.T) {
	asynClient := InitRedisMail()
	defer asynClient.Close()
	var err error

	for i := 0; i < 3; i++ {
		err = Emailer.SendPlainEmail([]string{"cuong@techmaster.vn"}, gofakeit.Sentence(10), gofakeit.Paragraph(3, 5, 7, "\n\n"))
		if err != nil {
			break
		}
	}
	assert := assert.New(t)
	assert.Nil(err)
}
