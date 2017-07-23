// 超时控制
package main

import (
	"fmt"
	"time"
)

func queryDb(ch chan int) {
	time.Sleep(time.Second)
	ch <- 100
}

func main() {
	ch := make(chan int)
	go queryDb(ch)
	t := time.NewTicker(time.Second)

	select {
	case v := <-ch:
		fmt.Println("result", v)
	case <-t.C:
		fmt.Println("timeout")
	}
}
