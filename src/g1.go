// g1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addBigInt(a, b string) (result string) {
	if len(a) == 0 && len(b) == 0 {
		return strconv.Itoa(0)
	}

	index1 := len(a) - 1
	index2 := len(b) - 1

	var left int
	var sum int

	fmt.Println(a, b)
	for index1 >= 0 && index2 >= 0 {
		c1 := a[index1] - '0'
		c2 := b[index2] - '0'

		sum = int(c1) + int(c2) + left

		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
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
}
