package test

import (
	"fmt"
	"gopgdemo/repo"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

func Test_exec_db(t *testing.T) {
	res, err := repo.DB.Exec(`CREATE TEMP TABLE test()`)
	assert.Nil(t, err)
	fmt.Println(res.RowsAffected())
}

/* chỉ cần đổi pgdb -> repo.DB là chạy được
https://pkg.go.dev/github.com/go-pg/pg/v10#Array
*/
func Test_array(t *testing.T) {
	src := []string{"one@example.com", "two@example.com"}
	var dst []string
	_, err := repo.DB.QueryOne(pg.Scan(pg.Array(&dst)), `SELECT ?`, pg.Array(src))
	assert.Nil(t, err)
	fmt.Println(dst)
}

//
func Test_hstore(t *testing.T) {
	src := map[string]string{"hello": "world", "foo": "bar"}
	var dst map[string]string
	_, err := repo.DB.QueryOne(pg.Scan(pg.Hstore(&dst)), `SELECT ?`, pg.Hstore(src))
	assert.Nil(t, err)
	fmt.Println(dst)
}

func Test_create_schema_table(t *testing.T) {

}
