package main

import (
	"fmt"
	"go_dev/day1/goroutine_example/calc"
	//	"../calc"
)

var p chan int

func main() {
	p = make(chan int, 100)
	go calc.Add(100, 391, p)
	go calc.Sub(132, 12, p)
	ss := <-p
	su := <-p
	fmt.Println("sum =", ss)
	fmt.Println("sub =", su)
}
