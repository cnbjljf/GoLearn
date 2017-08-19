// httpForm
package main

import (
	"fmt"
	"io"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar >
				<input type="text" name="in"/>
				<input type="text" name="in"/>
				<input type="submit" value="Submit"/>
				<from></body>
			</html>`

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "hello,world")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Tye", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}
}

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
		fmt.Println(err)
	}
}
