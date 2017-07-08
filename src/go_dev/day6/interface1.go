// interface1
package main

import (
	"fmt"
)

type carInterface interface {
	Run()
	getName() string
	diDI()
}

type byd struct {
	name string
}

func (b byd) getName() string {
	return b.name
}

func (b byd) Run() {
	fmt.Println("begin to run...", b.name)
}

func (b byd) diDI() {
	fmt.Println("DiDi ..", b.name)
}

type BMW struct {
	name string
}

func (b *BMW) getName() string {
	return b.name
}

func (b *BMW) Run() {
	fmt.Println("begin to run...", b.name)
}

func (b *BMW) diDI() {
	fmt.Println("DiDi ..", b.name)
}

func main() {
	var car carInterface
	bydcar := byd{"byd"}
	car = bydcar
	fmt.Println(car.getName())
	car.Run()
	car.diDI()
	fmt.Println("---------------------------------------")
	BMWcar := BMW{"BMW"}
	car = &BMWcar
	car.diDI()
	fmt.Println(car.getName())
	car.Run()

}
