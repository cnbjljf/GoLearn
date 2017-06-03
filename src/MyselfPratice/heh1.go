package main

import (
	"fmt"
)

type personInfo struct {
	ID   string
	Name string
}

func main() {
	for i := 0; i < 8; i++ {
		switch i {
		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(0)
		case 2:
			fmt.Println("fallthourgh")
			fallthrough
		case 4, 5, 6:
			fmt.Println("4,5,6")
		default:
			fmt.Println("default")
		}
	}
}
