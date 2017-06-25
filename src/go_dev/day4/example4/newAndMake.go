// newAndMake
package main

import (
	"fmt"
)

func main() {
	c := new([]int)

	*c = append(*c, 1, 2, 3, 4)
	fmt.Println(c)
}
