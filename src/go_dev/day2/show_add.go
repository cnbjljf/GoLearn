// show_add
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:输入一个数字就可以打印了")
		os.Exit(1)
	}

	a, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < a+1; i++ {
		sub_num := a - i
		fmt.Printf("%d+%d=%d\n", i, sub_num, a)
	}
}
