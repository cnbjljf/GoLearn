// template
package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Person struct { // 结构体名字必须都是大写的才可以识别在母板了
	Name  string
	Title string
	Age   string
}

func CreateHtml(w http.ResponseWriter) {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return "error"
	}

	p := Person{"marry", "my personal website", "11"}
	result := t.Execute(w, p)
	if result != nil {
		fmt.Println("there was an error:", err.Error())
		return "error"
}

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, CreateHtml(w))
}

func main() {
	http.HandleFunc("/test1", logPanic(SimpleServer))
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		fmt.Println(err)
	}
}

func logPanic(handle http.HandleFunc) http.HandleFunc {
	return func(write http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic:%v", request.RemoteAddr, x)
			}
		}()
		handle(write, request)
	}
}
