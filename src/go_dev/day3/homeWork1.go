// homeWork1
package main

import (
	"fmt"
	//	"strconv"
	"strings"
)

func nineNine() {
	// 99乘法表
	for i := 1; i < 10; i++ {
		x := i + 1
		for j := 1; j < x; j++ {
			s := i * j
			if j == i { // 如果J等于I，那么说明需要换行了，99乘法表里就是两个数相等的时候换行
				fmt.Printf("%dx%d=%d\n", j, i, s)
			} else {
				fmt.Printf("%dx%d=%d  ", j, i, s)
			}
		}
	}
}

func wanShu() {
	// 求1000以内的完全数
	// 什么是完全数：一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
	var wanShuList []int
	for i := 1; i < 1000; i++ {
		var ss []int
		for n := 1; n < i; n++ {
			if i%n == 0 && i != n { // 说明这个数是他的因数，能给整除,并且不能是自己本身
				ss = append(ss, n)
			}
		}
		total := 0
		for _, v := range ss { // 把所有的因数相加
			total += v
		}
		if total == i { // 判断因数之和是否等于这个数，是的话就是完全数
			wanShuList = append(wanShuList, i)
		}
	}
	fmt.Println(wanShuList)

}

func huiWen(a string) bool {
	// 判断其是否为回文。目前支持英文，支持中文,不支持中英混合与中数混合。
	//回文字符串是指从左到右读和从右到左读完全相同的字符串。
	var boolList []bool // 存放判断值，判断左右两侧是否一样
	var offset int      // 偏移量，for循环迭代下标的时候，中文的下标间隔是3，数字和英文的是1,所以要设置不同的下标偏移量
	var previous int    //存放第一个下标的值
	for i, _ := range a {
		if i == 0 {
			previous = i
		} else {
			if offset == 0 {
				offset = i - previous // 设置偏移量
			}
			index := len(a) - i - offset
			if byte(a[i]) == a[index] {
				boolList = append(boolList, true)
			} else {
				boolList = append(boolList, false)
			}
		}
	}
	for _, v := range boolList {
		if v == false {
			return false
		}
	}
	return true
}

func countNum(a string) map[string]int {
	// 输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数
	var countMap map[string]int
	countMap = make(map[string]int)
	for _, v := range a {
		countMap[string(v)] = strings.Count(a, string(v))
	}
	return countMap
}

func addBigInt(a, b string) {

}

func main() {
	//	nineNine()
	//	wanShu()
	//	fmt.Println(huiWen("上海自来水来自海上"))
	//	fmt.Println(countNum("asdfasdfc123123  aaaadoqpeiekc,gmntsdf9018234"))
	//	fmt.Println(addBigInt("22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222", "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"))
	//	a := "aaaaaa"
	//	fmt.Printf("%T\n", string(a[1]))
	//	fmt.Println(addBigInt("22222", "1111"))
	f := fmt.Sprintf("%d", 111)
	fmt.Println(f)
}
