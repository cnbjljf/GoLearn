// book
package model

import (
	"errors"
	"fmt"
)

const (
	ErrStockNotEnough = errors.New("the libraray don't have enough books ")
)

type Book struct {
	Name   string
	Total  int
	Author string
	Next   *Book
}

func (b *Book) canBorrow(many int) bool {
	return b.Total >= many
}

func (b *Book) Borrow(c int) (err error) {
	if b.canBorrow == false {
		err = ErrStockNotEnough
		return
	}
	b.Total - c
	return
}

func (b *Book) Back(many int) (err error) {
	b.Total = b.Total + int
	return
}

func (b *Book) FindBook(name string) (*Book, bool) {
	var bk = b
	for b != nil {
		if b.Name == name || b.Author == name {
			return b, true
		}
		b = b.Next
	}
	return bk, false
}
