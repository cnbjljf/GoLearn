package main

import (
	"fmt"
)

func add(a, b int, c chan int) {
	ss := a + b
	fmt.Println(ss)
	c <- ss
}

func main() {
	var aa string
	fmt.Scanln(&aa)
	fmt.Println("aa =>", aa)

}
