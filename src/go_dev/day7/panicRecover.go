// panicRecover
package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\r\n", e)
		}
	}()
	badCall()
	fmt.Println("after bad call")
}

func main() {
	test()
	fmt.Println("calling test")
}
