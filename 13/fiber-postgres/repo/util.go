package repo

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/* Sinh mã unique cho primary key
 */
func NewID() (id string) {
	id, _ = gonanoid.New(8)
	return id
}
