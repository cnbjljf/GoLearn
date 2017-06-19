// time1
package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间

	tt := time.Now().UnixNano()
	fmt.Println(time.Now())
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now().Minute())
	fmt.Println(time.Now().Month())
	fmt.Println(time.Now().Year())
	fmt.Println(time.Now().YearDay()) // 显示今天是今年的第几天
	fmt.Println(time.Now().Date())
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	now := time.Now()
	fmt.Println(now.Format("2006/1/02 15:04:05"))
	fmt.Println(now.Format("02/1/2006 15:04:05"))
	fmt.Println(now.Format("2006/1/02 15:04:05"))
	t2 := time.Now().UnixNano() - tt
	fmt.Println("how long ", t2/1000)

}
