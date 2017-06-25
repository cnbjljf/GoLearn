// qiboshu
package main

import (
	"fmt"
)

func fab(n int) {
	var a []uint64
	a = make([]uint64, n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-2] + a[i-1]
	}
	fmt.Println(a)
}

func main() {
	fab(10)
}
