// show_add
package main

import (
	"fmt"
)

func main() {
	a := 10
	for i := 0; i < a+1; i++ {
		sub_num := a - i
		fmt.Printf("%d+%d=%d\n", i, sub_num, a)
	}
}
