package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("hello", name, ",my name is ", u.Name)
}

func main() {
	u := User{1, "ok", 12}
	v := reflect.ValueOf(u)
	mv := v.MethodByName("hello")

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)
}
