// reflect_方法使用
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"student_name"`
	Age   int
	Score float32
	Sex   string
}

func (s student) print() {
	fmt.Println("--------")
	fmt.Println(s)
	fmt.Println("--------")
}

func (s student) set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect Struct")
		return
	}

	num := val.Elem().NumField()             // 有多少个字段
	val.Elem().Field(0).SetString("stu1000") // 对第一个字段设值
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n", num)
	tag := typ.Elem().Field(0).Tag.Get("json") // 获取结构体的tag，
	fmt.Printf("Tag = %s\n", tag)
	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)
	var params []reflect.Value
	val.Elem().Method(0).Call(params) // 调用第一个方法，并把参数params传入进去。
}

func main() {
	var stu = student{
		Name:  "stu01",
		Age:   18,
		Score: 99.8,
	}
	result, _ := json.Marshal(stu)
	fmt.Println("json result:", string(result))
	TestStruct(&stu)
	fmt.Println(stu)
}
