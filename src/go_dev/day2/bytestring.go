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

	he := fmt.Sprintf("%s %s", b, c)
	fmt.Println("Leo" < "haha", "lucy" == "jack")
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%c\n", a)
	fmt.Println(he)
}
