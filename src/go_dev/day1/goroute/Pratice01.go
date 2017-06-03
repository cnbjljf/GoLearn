package main

import "fmt"
import "time"

func calcNum(a int) {
	fmt.Println(a)
}

func main() {
	for i := 1; i < 100; i++ {
		go calcNum(i)
	}
	time.Sleep(time.Second)
}
