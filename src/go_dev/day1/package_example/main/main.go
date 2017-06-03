package main

import "fmt"
import "go_dev/day1/package_example/calc"

func main() {
	sum := calc.Add(100, 140)
	sub := calc.Sub(1400, 332)
	fmt.Println("sum=", sum)
	fmt.Println("sub =", sub)
}
