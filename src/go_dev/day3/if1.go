// if1
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputNum string
	inputNum = os.Args[1]
	f, err := strconv.Atoi(inputNum)
	if err != nil {
		fmt.Println("can not convert to int")
	} else {
		fmt.Printf("%d(%T)\n", f, f)
	}
}
