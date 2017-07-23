// chan的同步
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(10 * time.Second)
}

func sendData(ch chan string) {
	ch <- "first"
	ch <- "second"
	ch <- "thrid"
	ch <- "fourth"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}
