// book
package model

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day6/HomeWork/constValue"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrStockNotEnough       = errors.New("the libraray don't have enough books ")
	LoadOldDataError        = errors.New("load the book's data happend a error")
	NotFoundTheBookDataFile = errors.New("Not foudn the data file! please check the file of book's data")
	bookInfo                map[int]Book
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
		log.Fatalln("Saving book's data  happend a error:", err)
		return false
	}
	bkJsonData, _ := json.Marshal(bk)
	f.Write(bkJsonData)
	f.Sync()
	defer f.Close()
	fmt.Println("Saving book's data successfully!! ")
	//	for k, v := range bk {
	//		fmt.Printf("ID: %d, book:%v\n", k, v)
	//	}
	return true
}

func GetBookOldData() (bookInfo map[int]Book, err error) {
	// 获取图书所有信息
	bookInfo = make(map[int]Book)
	if constValue.Exist(constValue.BookDataFilePath) {
		oldF, err := ioutil.ReadFile(constValue.BookDataFilePath)
		if err != nil {
			log.Fatalln("happend a error when reading data from the book's file:", err)
			return bookInfo, LoadOldDataError
		}
		json.Unmarshal(oldF, &bookInfo)
		return bookInfo, nil
	}
	return bookInfo, NotFoundTheBookDataFile
}

func Delete(username, name string) bool {
	// 删除指定书籍的
	bookNewInfo := make(map[int]Book)
	bookInfo, _ := GetBookOldData()
	var i int
	for _, item := range bookInfo {
		if name == item.Name {
			constValue.Logger(username, "delete a book", "the book ["+name+"]was deleted!")
			continue
		} else {
			bookNewInfo[i] = item
			i++

		}

	}
	return SavebookJsonData(bookNewInfo)
}

func ManageBook(username string) {
	// 管理图书的，包含的管理动作有删除与修改
	for {
		bookInfo, _ := GetBookOldData()
		fmt.Println()
		for _, item := range bookInfo {
			fmt.Printf("name: %-8s ,author: %-8s ,stock: %-5d,published time: %-12s\n",
				item.Name, item.Author, item.Total, item.CreateTime)
		}
		fmt.Println("请输入书名,输入quit退出！")
		inputer := bufio.NewReader(os.Stdin)
		result, _, _ := inputer.ReadLine()
		bookName := strings.TrimSpace(string(result))

		if strings.ToLower(bookName) == "quit" {
			return
		}

		var flag bool
		for _, item := range bookInfo {
			if bookName == item.Name {
				flag = true
			}
		}
		if !flag {
			fmt.Println("无效的书名，没有找到这本书")
			continue
		}

		fmt.Println()
		fmt.Println(constValue.ActionChoice)

		inputer = bufio.NewReader(os.Stdin)
		result, _, _ = inputer.ReadLine()

		action, _ := strconv.Atoi(strings.TrimSpace(string(result)))

		switch action {
		case 1:
			Delete(username, bookName)
		case 2:
			fmt.Println("请输入新的信息,格式为作者，书籍总量，出版时间。必须以逗号分隔！")
			inputer := bufio.NewReader(os.Stdin)
			result, _, _ := inputer.ReadLine()
			allInfo := strings.Split(strings.TrimSpace(string(result)), ",")
			author := allInfo[0]
			total, _ := strconv.Atoi(allInfo[1])
			publishTime := allInfo[2]

			for k, item := range bookInfo {
				if bookName == item.Name {
					item.Author = author
					item.Total = total
					item.CreateTime = publishTime
					bookInfo[k] = item
					constValue.Logger(username, "modifing book's info", bookInfo[k])
					SavebookJsonData(bookInfo)
				}
			}

		}
	}

}
