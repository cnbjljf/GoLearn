// array
package main

import (
	"fmt"
)

func testArray() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{38, 283, 23, 2, 123, 4, 3, 2, 1, 1}
	var a3 = [...]int{1: 100, 3: 200}
	var a4 = [...]string{1: "hello", 2: "world"}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a)

}

func testArray2() {
	var a [2][5]int = [...][5]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}}
	for row, v := range a {
		for col, v1 := range v {
			fmt.Println("%d,%d = %d", row, col, v1)
		}
		fmt.Println()
	}
}

func main() {
	testArray2()
}
