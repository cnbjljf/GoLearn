// struct2
package main

import (
	"fmt"
	"math/rand"
)

type student struct {
	Name  string
	Age   int
	score float32
	next  *student
}

func showStruct(ss *student) {
	for ss != nil {
		fmt.Printf("%s,%d,%.1f,  %p\n", ss.Name, ss.Age, ss.score, ss.next)
		ss = ss.next
	}
}

func headInsert(p *student) *student {
	for i := 0; i < 5; i++ {
		stu := student{
			Name:  fmt.Sprintf("st%d", i),
			Age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}
		stu.next = p
		p = &stu
	}
	return p
}

func delNode(p *student, stuName string) *student {
	var tmp = p
	for p != nil {
		fmt.Println(p.Name, stuName)
		if p.Name == stuName {
			tmp.next = p.next
		}
		tmp = p
		p = p.next
	}
	return p
}

func rearInsert(p *student) {
	var tail = p // 赋值一个变量给tail，因为p是链表的第一个，
	// 也是最后一个，所以需要单独拿tail来操作
	for i := 0; i < 5; i++ {
		stu := student{
			Name:  fmt.Sprintf("st%d", i),
			Age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}
	//showStruct(p)
}

func main() {
	st0 := new(student)
	st0.Name = "Leo"
	st0.Age = 22
	st0.score = 90
	st0 = headInsert(st0)
	//	for st0 != nil {
	//		fmt.Printf("%s,%d,%.1f,  %p\n", st0.Name, st0.Age, st0.score, st0.next)
	//		st0 = st0.next
	//	}
	st0 = delNode(st0, "st2")
	showStruct(st0)
	//	fmt.Println()
	//	rearInsert(st0)
}
