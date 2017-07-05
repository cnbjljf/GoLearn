// DefineStruct
package main

import (
	"fmt"
)

type student struct {
	name  string
	age   int
	score float32
}

func main() {
	// 第一种申明方式
	stu := student{
		name:  "Leo",
		age:   19,
		score: 22.3,
	}
	fmt.Println(stu)

	// 第二种什么方式
	var stu1 = student{
		name:  "Leo",
		age:   19,
		score: 22.3,
	}
	fmt.Println(stu1)

	// 第三种
	stu3 := student{"Leo", 12, 33.4}
	fmt.Println(stu3)

	// 第四种
	stu4 := &student{"ljf", 22, 34.6}
	fmt.Println(stu4)

	// 第五种
	var stu5 *student = &student{"hehe", 56, 21.4}
	fmt.Println(stu5)
}
