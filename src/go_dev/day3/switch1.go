// switch1
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	randomInt := rand.Intn(100)
	fmt.Println("请输入数字")
	var inputNum int
	for {
		fmt.Scanln(&inputNum)
		switch {
		case randomInt > inputNum:
			fmt.Println("你输入的数字偏小，请重新输入")

		case randomInt < inputNum:
			fmt.Println("你输入的数字偏大，请重新输入")

		case randomInt == inputNum:
			fmt.Println("猜中了，幸运数字是", inputNum)
			os.Exit(1)
		default:
			fmt.Println()
		}
	}
}
