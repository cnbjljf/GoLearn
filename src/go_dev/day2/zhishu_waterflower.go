// zhishu
package main

import (
	"fmt"
	"strconv"
)

func zhishu() { // 统计100-200以内的素数

	var tmp []int // 用来存放最终素数结果的数组
	for i := 101; i < 200; i++ {
		var list []int            // 用来存放 i  被整除的 结果，然后统计这个结果的次数
		for a := 2; a < 10; a++ { // i被从2到9的数 除
			if i%a != 0 {
				list = append(list, i)
			}
		}

		if len(list) == 8 { // 长度等于8说明从2到9都没有被整除的数
			tmp = append(tmp, i)
		}
	}
	fmt.Println("素数", tmp)
}

func waterFlower() { // 统计100-999所有的水仙花数
	for i := 100; i < 1000; i++ {
		x := strconv.Itoa(i)                // 转为字符串格式，方便下一步切分
		x1, _ := strconv.Atoi(string(x[0])) // 取出这个数字的第一位数
		x2, _ := strconv.Atoi(string(x[1])) // 第二位
		x3, _ := strconv.Atoi(string(x[2])) // 第三位

		ret := x1*x1*x1 + x2*x2*x2 + x3*x3*x3
		if ret == i {
			fmt.Println("水仙花数", ret)
		}
	}
}

func nMulti(a int) { // 阶段次方之和
	ret := 0
	for i := 0; i <= a; i++ {
		ret += i * i
	}
	fmt.Printf("%d阶乘之和是%d\n", a, ret)
}

func main() {
	zhishu()
	waterFlower()
	nMulti(4)
}
