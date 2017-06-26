// addBigInt
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addBigInt(f1, f2 string) (result string) {
	if len(f1) == 0 && len(f2) == 0 {
		return string(0)
	}

	var index1 = len(f1) - 1 // the max index of the f1
	var index2 = len(f2) - 1 // the max index of the f2

	var left int

	for index1 >= 0 && index2 >= 0 { // 从右往左开始相加，加到最左边一位后跳出循环
		c1 := f1[index1] - '0' // 0的ASCII码是48,1是49等依次类推，所以得出来的int32类型的差
		c2 := f2[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' // 判断往前一位数进多少位,用%求的是余数，
		fmt.Printf("%c\n", c3)
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	//走下面两个循环的的话，说明其中有个数是大于另一个数的，那么就需要单独处理下。
	for index1 >= 0 { // 进行最左边数的相加,
		c1 := f1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' // 判断往前一位数进多少位,用%求的是余数
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c2 := f2[index2] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' // 判断往前一位数进多少位,用%求的是余数
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println("read error from stdin:", err)
		return
	}

	strSlice := strings.Split(string(result), "+")
	if len(strSlice) != 2 {
		fmt.Println("please input a+b")
		return
	}
	f1 := strings.TrimSpace(strSlice[0])
	f2 := strings.TrimSpace(strSlice[1])

	fmt.Println(addBigInt(f1, f2))
}
