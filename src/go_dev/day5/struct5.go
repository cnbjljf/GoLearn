// struct5
package main

import (
	"fmt"
)

type student struct {
	name  string
	age   int
	sex   string
	score float32
}

func (p *student) init(name string, age int, sex string, score float32) {
	p.name = name
	p.age = age
	p.sex = sex
	p.score = score
}

func (p *student) get() *student {
	fmt.Println(p)
	return p
}

func main() {
	var stu student
	stu.init("leo", 22, "man", 89.5)
	stu.get()
}
