// func1
package main

import (
	"fmt"
)

type add_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op add_func, a, b int) int {
	return op(a, b)
}

func returnName(a, b int) (total, sub int) {
	total = a + b
	sub = a / b
	return
}

func add2(arg ...int) int {
	s := 0
	for _, v := range arg {
		s += v
	}
	return s
}

func add3(a, b int, args ...int) int {
	s := a + b
	for _, v := range args {
		s += v
	}
	return s
}

func addStr(a string, args ...string) string {
	s := a
	for _, v := range args {
		s += v
	}
	return s
}

func main() {
	c := add
	fmt.Println(c)

	su := c(10, 20)
	fmt.Println(su)

	ss := operator(add, 12, 30)
	fmt.Println(ss)

	t, s := returnName(10, 34)
	fmt.Println(t, s)

	fmt.Println(add2(1, 2, 3, 4, 5, 6, 10))
	fmt.Println(add3(2, 3, 111, 333, 444, 555))
	fmt.Println(addStr("asdf", "adfuwe", "123", "siw", "sdc"))
}
