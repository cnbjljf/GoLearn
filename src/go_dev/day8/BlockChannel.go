// BlockChannel
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	time.Sleep(10 * time.Second)
}

func sendData(ch chan string) {
	var i int
	for {
		var str string
		str = fmt.Sprintf("stu %d", i)
		fmt.Println("write:", str)
		ch <- str
		i++
	}
}
