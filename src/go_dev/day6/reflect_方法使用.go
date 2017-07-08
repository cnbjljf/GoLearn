// reflect_方法使用
package main

import (
	"fmt"
)

type student struct {
	Name  string `json:"student_name"`
	Age   int
	Score float32
	Sex   string
}

func (s Student) set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
}

func main() {
}
