// print&
package main

import (
	"fmt"
)

func main() {
	var a int = 10
	fmt.Println(&a)

	var p *int
	p = &a
	fmt.Println("before =", *p)
	*p = 100
	fmt.Println("after =", a)

	var b int = 999
	p = &b
	*p = 4
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("14%10", 4%10)
	fmt.Println('1' - '0')
}
