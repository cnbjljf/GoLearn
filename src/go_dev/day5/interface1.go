// interface1
package main

import (
	"fmt"
)

type test interface {
	Fuck()
	Sleep()
}

type people struct {
	name string
	age  int
	sex  string
}

func (p people) Fuck() {
	fmt.Println("fuck you!!", p.name)
}

func (p people) Sleep() {
	fmt.Println("i am sleep now", p.name)
}

func main() {
	var t test
	var st = people{
		name: "Leo",
		age:  19,
		sex:  "asdf",
	}

	t = st
	t.Fuck()
	t.Sleep()
}
