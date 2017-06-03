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

func (u User) Hello() {

}

func main() {
	u := User{1, "ok", 12}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name1") // FieldByName 是获取指定字段的方法
	if !f.IsValid() {           // 判断是否有效，前面添加了！，表示取反
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String { // 如果f的类型是反射率IM的字符串，那么就通过SetString来设置
		f.SetString("BYEBYE")
	}
}
