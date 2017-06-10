// 16to32
package main

import (
	"fmt"
)

func main() {
	var n int16 = 23
	var m int32
	var a byte
	a = 'a'
	fmt.Printf("%c\n", a)
	//	m = n   类型不一样，所以不能够直接等于号赋值
	m = int32(n)
	fmt.Println(m, n)
}
