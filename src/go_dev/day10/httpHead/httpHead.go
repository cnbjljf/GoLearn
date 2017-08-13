// httpHead
package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
}

func main() {
	for _, v := range url {
		c := http.Client{ // 指定http的超时时间，
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					timeout := time.Second * 2 // 这里设置超时时间为2秒
					return net.DialTimeout(network, addr, timeout)
				},
			},
		}

		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s failed,err:", v, err)
			continue
		}
		fmt.Printf("head succ,status:%v\n", resp.Status)

	}
}
