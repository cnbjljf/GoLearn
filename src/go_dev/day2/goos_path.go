// goos_path
package main

import (
	"fmt"
)

func main() {
	//	fmt.Println("GOOS", os.Getenv("GOOS"))
	//	fmt.Println("path", os.Getenv("PATH"))
	p := make(chan int, 10)

	p <- 1
	p <- 2
	p <- 3
	close(p)
	for i := range p {
		fmt.Println(i)
	}
	//	c := 10
	//	fmt.Println("c", c)
	//	fmt.Println("P", p)
	a := 3
	b := 4
	b, a = a, b
	fmt.Println(a, b)
}
