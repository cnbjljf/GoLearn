// bytestring
package main

import (
	"fmt"
)

func main() {
	a := 'b'
	b := `hehehe \n \nasdfasdfasdf
	asdfasdfasfjjkl  \t \n asdfasdfasf`

	c := "aaaa"
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%c\n", a)
}
