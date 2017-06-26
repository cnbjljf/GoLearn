// g1
package main

import (
	"fmt"
	"strings"
)

func main() {

	aa := "14+12"
	bb := strings.Split(aa, "+")
	f1 := []rune(bb[0])
	f2 := []rune(bb[1])

	lenf1 := len(f1)
	lenf2 := len(f2)
	var lenMax int
	//	var lenMin int

	left := 0

	if lenf1 >= lenf2 {
		//		lenMin = lenf2
		lenMax = lenf1
	} else {
		//		lenMin = lenf1
		lenMax = lenf2
	}
	for i := 1; i <= lenMax; i++ {
		index := lenMax - i
		sum := (f1[index] - '0') + (f2[index] - '0') + rune(left)
		fmt.Println(sum)
		if sum > 10 {
			left = 1
		}
	}

	fmt.Printf("%c %c", f1, f2)
}
