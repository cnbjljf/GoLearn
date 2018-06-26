// mysql_select
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123..aa@tcp(192.168.56.14:3306)/test")
	if err != nil {
		fmt.Println("connect mysql failed", err)
		return
	}
	Db = database
}

func main() {
	var person []Person
	err := Db.Select(&person, "select * from person;")
	if err != nil {
		fmt.Println("exec failded", err)
		return
	}
	fmt.Println("select succ:", person)
}
