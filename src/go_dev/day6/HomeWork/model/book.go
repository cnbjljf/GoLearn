// book
package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day6/HomeWork/constValue"
	"log"
	"os"
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

func SavebookJsonData(bk map[int]Book) bool {
	// 保存图书信息到文本文件的
	f, err := os.Create(constValue.BookDataFilePath)
	if err != nil {
		log.Fatalln("Saving data  happend a error:", err)
		return false
	}
	bkJsonData, _ := json.Marshal(bk)
	f.Write(bkJsonData)
	f.Sync()
	defer f.Close()
	fmt.Println("Saving data successfully!! these book infomation is :")
	for k, v := range bk {
		fmt.Printf("ID: %d, book:%v\n", k, v)
	}
	return true
}
