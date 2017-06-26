// p1
package main

import (
	"fmt"
	"time"
)

var d int

const (
	man    = 1
	female = 2
)

func main() {
	s := time.Now().Unix()
	//	fmt.Println(man, s)
	fmt.Println(s, s%female)
	fmt.Println(s / female)
	if s%female == 0 {
		fmt.Println("整除", female)
	} else {
		fmt.Println("不能整除", man)
	}
}
