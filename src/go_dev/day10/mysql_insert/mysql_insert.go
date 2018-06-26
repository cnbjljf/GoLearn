// mysql_insert
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	//  所有字段必须和数据库的保持一致，如果第一列使用了和数据库字段不一样的名字，那么就需要在
	//  ·· 之间写明真正的字段名
	UserId   int    `db:"user_id`
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
	// 连接格式为(数据库类型，用户名:密码@tcp(ip:port)/库名
	database, err := sqlx.Open("mysql", "root:123..aa@tcp(192.168.56.14:3306)/test")
	if err != nil {
		fmt.Println("open mysql faild,", err)
		return
	}
	Db = database
}

func main() {
	r, err := Db.Exec("insert into person values(?,?,?,?)", 2, "Leo", "man", "test@qq.com")
	if err != nil {
		fmt.Println("exec sql command failed", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("get the insert ID failed", err)
		return
	}
	fmt.Println("insert successful, the id :", id)
}
