// digui
package main

import (
	"fmt"
	"time"
)

func recusive(n int) {
	fmt.Println("hello", n)
	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	recusive(n + 1)
}

func factor(n int) int {
	if n == 1 {
		return 1
	}
	fmt.Println("factor n :", n)
	return factor(n-1) * n
}

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	fmt.Println("fab n :", n)
	return fab(n-1) + fab(n-2)
}

func main() {
	//	recusive(1)
	factor(3)
	fab(4)
}
