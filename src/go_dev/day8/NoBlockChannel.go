// BlockChannel
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 10) // 唯一的区别在于这里，有10 个插入字符串次数作为缓冲区
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
