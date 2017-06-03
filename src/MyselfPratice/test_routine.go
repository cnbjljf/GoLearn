package main

import (
	"fmt"
)

var p chan int

func calcNum(a int, b int, p chan int) {
	rt := a + b
	p <- rt
}

func main() {
	p = make(chan int, 200)
	for i := 0; i < 100; i++ {
		go calcNum(i, i+12, p)
	}
	for {
		if len(p) == 0 {
			break
		}
		tt := <-p
		fmt.Println(tt)

	}
}
