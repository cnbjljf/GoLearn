// reverserStr
package main

import (
	"fmt"
)

func main() {
	a := "Hello World!"
	lenA := len(a)
	var b string
	for i := 1; i <= lenA; i++ {
		b = string(b) + string(a[lenA-i])
	}
	fmt.Printf("%s", b)
}
