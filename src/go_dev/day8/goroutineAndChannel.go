// goroutineAndChannel
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(100 * time.Second)
}

func sendData(ch chan string) {
	ch <- "asdfasd"
	ch <- "second"
	ch <- "third"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}
