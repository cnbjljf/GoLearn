// http_simple
package main

import (
	"fmt"
	"net/http"
)

func Hellow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle Hello")
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", Hellow)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
