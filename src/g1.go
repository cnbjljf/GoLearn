// g1
package main

import (
	"bufio"
	"fmt"
<<<<<<< HEAD
	"os"
=======
>>>>>>> 1e76b4995291dd3a83c0d5ea17def2da75f22559
	"strconv"
	"strings"
)

func addBigInt(a, b string) (result string) {
	if len(a) == 0 && len(b) == 0 {
		return strconv.Itoa(0)
	}

<<<<<<< HEAD
	index1 := len(a) - 1
	index2 := len(b) - 1

	var left int
	var sum int

	fmt.Println(a, b)
	for index1 >= 0 && index2 >= 0 {
		c1 := a[index1] - '0'
		c2 := b[index2] - '0'

		sum = int(c1) + int(c2) + left

=======
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
>>>>>>> 1e76b4995291dd3a83c0d5ea17def2da75f22559
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
<<<<<<< HEAD
		c3 := (sum % 10) - '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ss, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	strSlice := strings.Split(string(ss), "+")
	if len(strSlice) != 2 {
		fmt.Println("please input a+b")
		return
	}

	a := strSlice[0]
	b := strSlice[1]
	fmt.Println(addBigInt(a, b))
=======
		rt = s1 + rt
	}

	fmt.Println(rt)
>>>>>>> 1e76b4995291dd3a83c0d5ea17def2da75f22559
}
