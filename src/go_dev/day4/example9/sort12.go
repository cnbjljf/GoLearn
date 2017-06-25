// sort12
package main

import (
	"fmt"
	"sort"
)

func sortStr() {
	var a = [...]string{"abc", "efg", "b", "b", "cc", "eeafc"}
	sort.Strings(a[:])
	fmt.Println(a)
}

func sortInt() {
	var a = [...]int{102, 2, 8, 1, 3}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2222.3, 5.3, 3.9034, 1003.111}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1222, 345, 39, 141}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 39)
	fmt.Println(index)
}

func main() {
	//	sortStr()
	//	sortInt()
	//	testFloat()
	testIntSearch()
}
