//  反射typekind
package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b) // 获取类型
	fmt.Println(t)

	v := reflect.ValueOf(b) // 获取值
	k := v.Kind()           // 获取数据类型
	fmt.Println(k)

	iv := v.Interface()     // 转为interface类型
	stu, ok := iv.(student) // 转为student类型
	if ok {                 // 转换成功的话
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	c := val.Int() // 转为int64类型
	fmt.Printf("get value from interface{} %d\n", c)
}

func main() {
	var a = student{"stu01", 18, 99}
	test(a)
	testInt(123)
}
