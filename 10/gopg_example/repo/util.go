package repo

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

/* Sinh mã unique cho primary key
 */
func NewID() (id string) {
	id, _ = gonanoid.New(8)
	return
}

/* Kiểm tra xem biến a có nằm trong mảng int_arr
 */
func int_in_array(a int, int_arr []int) bool {
	for _, b := range int_arr {
		if b == a {
			return true
		}
	}
	return false
}
