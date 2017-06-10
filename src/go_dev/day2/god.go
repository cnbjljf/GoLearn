// god
package main

import (
	"fmt"
)

var g string = "G"

func n() {
	fmt.Println(g)
}

func m() {
	g := "O"
	fmt.Println(g)
}

func main() {
	n()
	m()
	n()
}
