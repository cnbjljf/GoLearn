// interface2_扩展sort
// 此次代码流程参考连接：https://golang.org/pkg/sort/#Interface
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type student struct {
	name string
	id   string
	age  int
}

type stuArray []student

func (st stuArray) Len() int {
	return len(st)
}

func (st stuArray) Less(i, j int) bool {
	return st[i].name > st[j].name
}

func (st stuArray) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
}

func main() {
	fmt.Println("Hello World!")
	var stus stuArray
	for i := 0; i < 10; i++ {
		stu := student{
			fmt.Sprintf("stu%d", rand.Intn(100)),
			fmt.Sprintf("110%d", rand.Int()),
			rand.Intn(70)}
		stus = append(stus, stu)
	}
	for _, v := range stus {
		fmt.Println(v)
	}
	fmt.Println()
	fmt.Println()
	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}

}
