// mysql_delete
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	conn, err := sqlx.Open("mysql", "root:123..aa@tcp(192.168.56.14:3306)/test")
	if err != nil {
		fmt.Println("connect mysql failed,e:", err)
		return
	}
	Db = conn
}

func main() {
	result, err := Db.Exec("delete from person  where user_id=?", 2)
	if err != nil {
		fmt.Println("execute sql command error", err)
		return
	}

	fmt.Println("result", result)

}
