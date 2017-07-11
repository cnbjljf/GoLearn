// interface到具体类型的转换
package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func test(a interface{}) {
	b, ok := a.(int) //把interface转为int类型
	if ok == false {
		fmt.Println("conver faild")
		return
	}
	fmt.Println(b)
}

// 通过switch 语句来匹配
func test2(a ...interface{}) {
	for i, v := range a {
		switch v.(type) { // type 是解释器会自动赋值的
		case string:
			fmt.Printf("the index [%d] is [%T]\n", i, v)
		case float32, float64:
			fmt.Printf("the index [%d] is [%T]\n", i, v)
		case int, int32, int64, int8, int16:
			fmt.Printf("the index [%d] is [%T]\n", i, v)
		case student:
			fmt.Printf("the index [%d] is [%T]\n", i, v)
		case *student:
			fmt.Printf("the index [%d] is [%T]\n", i, v)
		}
	}
}

func main() {
	var c student
	test(c)

	var b int
	b = 10
	test(b)
	test2(1, 3.322, "fuck man", c, &c)
}
