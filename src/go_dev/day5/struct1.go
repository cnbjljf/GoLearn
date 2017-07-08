// struct1
package main

import (
	"fmt"
)

type student struct {
	Name  string
	Age   int
	score float32
}

func main() {
	// 常见对struct赋值的方法
	var st1 student
	st1.Age = 20
	st1.Name = "Leo"
	st1.score = 99.5
	fmt.Println(st1)

	// var 定义的时候赋值,st2 为指针类型,也就是引用类型
	var st2 *student = &student{
		Name:  "Leo",
		Age:   22,
		score: 22.6,
	}
	fmt.Println(st2)

	// var 定义的时候赋值，也就是值类型
	var st3 = student{
		Name:  "ljf",
		Age:   22,
		score: 36.3,
	}
	fmt.Println(st3)
}
