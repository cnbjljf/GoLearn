// 定时器的使用
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("hello,", v)
		f := time.NewTicker(3 * time.Second)
		for fv := range f.C {
			fmt.Println("second --->,", fv)
		}
	}

}
