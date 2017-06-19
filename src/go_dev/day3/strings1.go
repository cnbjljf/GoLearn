// strings1
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func strIndex(str, substr string) (int, int) {
	//  写一个函数返回一个字符串在另一个字符串的首次出现和最后出现位置
	firstIndex := strings.Index(str, substr)
	lastIndex := strings.LastIndex(str, substr)
	return firstIndex, lastIndex
}

func main() {
	url := "www.baidu.com"
	url2 := "http://bwwwm.baidu.com"

	/* 判断一个url是否以http://开头，如果不是，则加上http:// */
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
		fmt.Println(url)
	}

	if strings.HasPrefix(url2, "http") {
		fmt.Println(url2)
	}

	//  判断一个路径是否以“/”结尾，如果不是，则加上/。

	p1 := "/var/www"
	p2 := "/var/www/"

	if !strings.HasSuffix(p1, "/") {
		p1 = p1 + "/"
		fmt.Println(p1)
	}

	if strings.HasSuffix(p2, "/") {
		fmt.Println(p2)
	}

	// 判断str在s中首次出现的位,置如果没有 出现，则返回-1
	fmt.Println(len(url2), strings.Index(url2, "b"))

	// 判断str在s中最后出现的位置，如果没有出现，则返回-1
	fmt.Println(strings.LastIndex(url2, "m"))

	fmt.Println(strIndex("www.baidu.com and pan.baidu.com", "baidu"))

	//  字符串替换,最后一个参数２表示只替换２次，如果是－１那么就是全部替换
	fmt.Println(strings.Replace("www.baidu.com,pan.baidu.com,baidu.cn", "baidu", "sina", 2))

	//字符串计数
	fmt.Println("字符串计数", strings.Count("www.baidu.com,pan.baidu.com,baidu.cn", "baidu"))

	// 转为小写
	fmt.Println(strings.ToLower("转为小写 WHAT THE HELL?MAN"))
	// 转为大写
	fmt.Println("转为大写", strings.ToUpper(url2))

	// 去掉字符串首尾空白字符
	fmt.Println(strings.TrimSpace("   wwwbaidu.com   "))

	// 去掉字符串首尾cut字符
	fmt.Println(strings.Trim("www.baidu.com.www", "www"))

	//去掉字符串首cut字符
	fmt.Println(strings.TrimLeft("www.baidu.com.www", "www"))

	//去掉字符串首cut字符
	fmt.Println(strings.TrimRight("www.baidu.com.www", "www"))

	// 返回str空格分隔的所有子串的slice
	tt1 := strings.Fields("www. baidu.co m.w ww")
	fmt.Println(tt1, len(tt1))

	// 返回str split分隔的所有子串的slice
	tt := strings.Split("www.baidu.com.www", ".")
	fmt.Printf("%T %v\n", tt, tt)

	// 用sep把s1中的所有元素链接起来
	var ss []string
	ss = []string{"aa", "bb", "cc"}
	fmt.Println(strings.Join(ss, "-"))

	// 把一个整数i转成字符串
	ss1 := strconv.Itoa(10)
	fmt.Printf("%T %s\n", ss1, ss1)

	// 把一个字符串转成整数
	ss2, _ := strconv.Atoi("10")
	fmt.Printf("%T %v\n", ss2, ss2)

}
