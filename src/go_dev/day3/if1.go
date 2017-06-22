// if1
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputNum string
	fmt.Println("= 请输入需要转换数字")
	//	fmt.Scanln(&inputNum)
	fmt.Scanf("%s", &inputNum)
	f, err := strconv.Atoi(inputNum)
	if err != nil {
		fmt.Println("can not convert to int")
	} else {
		fmt.Printf("%d(%T)\n", f, f)
	}
}
