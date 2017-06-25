// bibao1
package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func bibao(suffix string) func(name string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}

func main() {
	f1 := bibao(".txt")
	fmt.Println(f1("name.txt"))
	fmt.Println(f1("hehe"))

	f2 := adder()
	fmt.Println(f2(10))
	fmt.Println(f2(12))
}
