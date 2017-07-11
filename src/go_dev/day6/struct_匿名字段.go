// struct_匿名字段
package main

import (
	"fmt"
)

type car struct {
	Name  string
	brand string
}

type car2 struct {
	price int
	brand
}

type car3 struct {
}

func main() {

}
