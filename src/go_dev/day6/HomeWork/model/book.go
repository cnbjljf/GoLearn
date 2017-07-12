// book
package model

import (
	"errors"
)

var (
	ErrStockNotEnough = errors.New("the libraray don't have enough books ")
)

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func (b *Book) canBorrow(many int) bool {
	return b.Total >= many
}

func (b *Book) Borrow(c int) (err error) {
	if b.canBorrow(c) == false {
		err = ErrStockNotEnough
		return
	}
	b.Total = b.Total - c
	return
}

func (b *Book) Back(many int) (err error) {
	b.Total = b.Total + many
	return
}

//func (b *Book) FindBook(bkjson map[int]book.Book, name string) (*Book, bool) {

//	for b != nil {
//		if b.Name == name || b.Author == name {
//			return b, true
//		}
//	}
//	return bk, false
//}
