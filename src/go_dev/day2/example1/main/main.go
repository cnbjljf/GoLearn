// main
package main

import (
	"fmt"
	t "go_dev/day2/example1/add"
)

func init() {
	fmt.Println("begin to init")
}

func main() {
	fmt.Println("Name:", t.Name)
	fmt.Println("Age:", t.Age)
}
