// httptemplate
package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  string // struct字段必须都是大写的，这样模版才可以调用到
}

func main() {
	t, err := template.ParseFiles("H:/Go/GoLearn/src/go_dev/day10/httpTemplate/if.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", Age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
