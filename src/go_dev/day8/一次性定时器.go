// 一次性定时器
package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case <-time.After(time.Second):
		fmt.Println("after")
	}
}
