// g1
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	aa := "99+1"
	bb := strings.Split(aa, "+")
	f1 := []rune(bb[0])
	f2 := []rune(bb[1])

	lenf1 := len(f1) //2
	lenf2 := len(f2) // 1
	var lenMax int
	//	var lenMin int

	left := 0
	var rt string
	var sum rune

	if lenf1 >= lenf2 {
		//		lenMin = lenf2
		lenMax = lenf1
	} else {
		//		lenMin = lenf1
		lenMax = lenf2
	}
	for i := 1; i <= lenMax; i++ {
		index := lenMax - i
		if index >= lenf1 {
			if index < lenf2 {
				sum = (f2[index] - '0') + rune(left)
			}
		} else if index >= lenf2 {
			if index < lenf1 {
				sum = (f1[index] - '0') + rune(left)
			}

		} else {
			sum = (f1[index] - '0') + (f2[index] - '0') + rune(left)
		}
		s1 := strconv.Itoa(int(sum))
		fmt.Printf("%T \t %v\n", sum, sum)
		if sum >= 10 {
			left = 1
		}
		rt = s1 + rt
	}

	fmt.Println(rt)
}
