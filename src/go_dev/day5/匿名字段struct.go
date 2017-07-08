// 匿名字段struct
package main

import (
	"fmt"
)

type car struct {
	name  string
	brand string
}

type car2 struct {
	brand string
}

type allCar struct {
	car
}

func main() {
	fmt.Println("Hello World!")
}
