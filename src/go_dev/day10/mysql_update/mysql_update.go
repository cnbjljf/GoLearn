// mysql_update
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
		fmt.Println("connect mysql failed,", err)
		return
	}
	Db = conn
}

func main() {
	_, err := Db.Exec("update  person set username=? where user_id=?", "Yq", 1)
	if err != nil {
		fmt.Println("execute sql command happend a error,", err)
		return
	}
	fmt.Println("update successfully")
}
