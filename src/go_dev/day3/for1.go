// for1
package main

import (
	"fmt"
	//	"time"
)

func main() {
	aa := "A"
	for i := 0; i <= 5; i++ {
		fmt.Println(aa)
		aa = aa + "A"
	}

	str := "hello world,中国"
	for i, v := range str {
		if i < 1 {
			fmt.Printf("index:[%d]  val[%c]  \n", i, v)
		} else if i > 5 {
			break
		}
	}

	xx := 0
	for xx < 10 {
		xx++
		fmt.Println("xx", xx)
	}

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is :%d and j is :%d\n", i, j)
		}
	}

	for i := 0; i < 10; i++ {
		if i > 3 {
			fmt.Println(i)
			goto LABEL2
		}
	}
LABEL2:
	fmt.Println("break ")

}
