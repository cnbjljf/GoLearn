// 接口式通用链表
package main

import (
	"fmt"
)

func main() {
	var link Link
	for i := 0; i < 10; i++ {
		link.InsertHead(fmt.Sprintf("%d \n", i))
	}

}
