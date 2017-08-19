// accessWebsite
package main

import (
	"fmt"
	"net/http"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {
	for _, v := range url {
		res, _ := http.Head(v)
		fmt.Println("head success,status:%s\n", res.Status)
	}
}
