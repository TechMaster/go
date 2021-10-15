package util

import (
	"math/rand"
	"strconv"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

/* Sinh mã unique cho primary key
 */
func NewID(length ...int) (id string) {
	if len(length) == 0 {
		id, _ = gonanoid.New(8)
	} else {
		id, _ = gonanoid.New(length[0])
	}
	return
}

func RandomId() (id string) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000)
	id = strconv.Itoa(randomNumber)
	return id
}