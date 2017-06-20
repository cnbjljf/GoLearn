// func2
package main

import (
	"fmt"
)

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}

func main() {
	//	a()
	fmt.Println(4 / 2)
}
