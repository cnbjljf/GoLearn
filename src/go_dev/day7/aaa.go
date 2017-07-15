// aaa
package main

import (
	"fmt"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	var str = "stru01 18 89 92"
	var stu student
	fmt.Sscanf(str, "%d %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}
