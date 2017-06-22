// 111
package main

import (
	"fmt"
)

result = func(a,b int) int{
	return a+b
}

func test(a,b int) int{
	return result(a,b)
}
func main() {
	fmt.Println(test(10,3))
}
