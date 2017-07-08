// refelect_几个方法
package main

import (
	"fmt"
	"reflect"
)

type unknowType struct {
	s1 string
	s2 string
	s3 string
}

func (n unknowType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = unknowType{"ada", "Go", "obera"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)

	kd := value.Kind()
	fmt.Println("kind", kd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d:%v\n", i, value.Field(i))
	}
	results := value.Method(0).Call(nil)
	fmt.Println(results)

}
