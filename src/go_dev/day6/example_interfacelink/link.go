// link
package main

import (
	"fmt"
)

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	// 头部插入 链表，意味着最近插入的数据是最靠前的。先入先出
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.tail == nil {
		p.tail = data
		p.next = data
		return
	}
	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {
	// 尾部插入，意味着最近插入的数据都在尾部，先入后出
	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = data
		p.head = data
	}
	p.tail.next = node
	p.tail = node
}

func (p *Link) Trans() {
	// 打印出这个链表
	q := p.head
	for q != nil {
		fmt.Println(q.data)
		q = q.next
	}
}
