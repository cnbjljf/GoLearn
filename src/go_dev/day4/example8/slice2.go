// slice2
package main

import (
	"fmt"
)

func testSlice() {
	var a [5]int = [...]int{1, 2, 3, 4, 5}
	s := a[1:]
	fmt.Println("before len[%d] cap[%d]\n", len(s), cap(s))
	s[1] = 100
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("before a:", a)

	s = append(s, 10)
	s = append(s, 10)
	fmt.Println("after len[%d] cap[%d]\n", len(s), cap(s))
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s[1] = 1000
	fmt.Println("after a:", a)
	fmt.Println(s)
	fmt.Println("s=%p a[1]=%p\n", s, &a[1])
}

func main() {
	fmt.Println("Hello World!")
}
