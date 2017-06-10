// randnum
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var a []int
	var b []float64
	for i := 0; i < 11; i++ {
		a = append(a, rand.Intn(100))
		b = append(b, rand.Float64()*10)
	}
	fmt.Println(a)
	fmt.Println(b)
}
