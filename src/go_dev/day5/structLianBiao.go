// structLianBiao
// struct 链表
//
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type student struct {
	name  string
	age   int
	score float32
	next  *student
}

func showStruct(p *student) {
	for p != nil {
		fmt.Println(p)
		p = p.next
		time.Sleep(20 * time.Millisecond)
	}
}

func headInsert(p *student) *student {
	// 头部插入结构体，也就是在for循环里面越后插入的结构体越在前面
	for i := 0; i < 6; i++ {
		stu := student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}

		stu.next = p
		p = &stu
	}
	//	showStruct(p)
	return p
}

func tailInsert(p *student) student {
	var tail = p

	// 尾部插入，也就是在for循环里面越在后面插入的值，越在最后
	for i := 0; i < 6; i++ {
		stu := student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
	return *p
}

func delNode(s *student, name string) student {
	var ss = s
	previous := s
	for s != nil {
		if s.name == name {
			previous.next = s.next
		}
		previous = s
		s = s.next
	}
	return *ss
}

func main() {
	var st0 = student{
		name:  "Leo",
		age:   22,
		score: 99.5,
	}
	//	st1 := *headInsert(&st0)
	//	showStruct(&st1)
	fmt.Println("尾部插入")
	st2 := tailInsert(&st0)
	st2 = delNode(&st2, "stu3")
	showStruct(&st2)
}
