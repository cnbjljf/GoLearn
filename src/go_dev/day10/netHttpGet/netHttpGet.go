// netHttpGet
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1")
	if err != nil {
		fmt.Println("happend a error!", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("get data err :", err)
		return
	}

	fmt.Println(string(data))
}
