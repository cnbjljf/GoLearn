// main
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_dev/day6/HomeWork/constValue" // 定义一些常量，相当于配置文件一样
	"go_dev/day6/HomeWork/model"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	StockNum = 10000 // 图书馆图书容量
)

var bookJson map[int]model.Book
var stuJson map[int]model.Student

func AddBook() {

	// 增加书籍的，录入书籍的信息的格式必须是  书名,作者,出版时间,多少本   必须以逗号分隔。
	// 先读取文件里买的数据，然后在写入，避免覆盖之前的数据
	bookJson = make(map[int]model.Book)

	bookJson, err := model.GetStuOldData() //先判断文件是否存在，存在才读取老文件
	if err != nil {
		log.Fatalln(err)
		return
	}

	var i = 0

	for i < StockNum { // 这个循环来标记出我们存入到文件的数据现在到多少本了,不能超过库存容量
		_, ok := bookJson[i]
		if ok == false {
			break
		}
		i++
	}

	for i < StockNum {
		fmt.Println("Please Input your book's infoenter quit will quit this func, and the format is ")
		fmt.Println(" name, author ,published , how many")
		fmt.Println("举个例子：    aaa,Leo,20170705,12 , 必须以逗号分隔")

		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happed a error:", err)
			return
		}
		rawData := strings.Split(string(result), ",")
		bookName := strings.TrimSpace(rawData[0])
		if bookName == "quit" {
			fmt.Println("exit to the add Book")
			break
		}
		if len(rawData) != 4 {
			fmt.Println("you weren't provide some right arguments")
			continue
		}
		author := strings.TrimSpace(rawData[1])
		published := strings.TrimSpace(rawData[2])
		many, _ := strconv.Atoi(strings.TrimSpace(rawData[3]))
		if bookName == "" && author == "" && published == "" && many == 0 {
			continue
		}
		bookInfo := model.Book{
			Name:       bookName,
			Total:      many,
			Author:     author,
			CreateTime: published,
		}
		bookJson[i] = bookInfo
		i++
	}
	if model.SavebookJsonData(bookJson) {
		return
	} else {
		fmt.Println("Save data happend a error!")
	}
}

func AddStu() {
	// 添加学生信息的，录入学生的信息必须是i
	stuJson = make(map[int]model.Student)

	stuJson, err := model.GetStuOldData() //先判断文件是否存在，存在才读取老文件
	if err != nil {
		log.Fatalln(err)
		return
	}

	var i = 0
	for { // 这个循环来标记出我们存入到文件的数据现在到多少本了,不能超过库存容量
		_, ok := stuJson[i]
		if ok == false {
			break
		}
		i++
	}

	for {
		fmt.Println("Please Input your student's infomation,enter quit will quit this func, and the format is ")
		fmt.Println(" name, ID ,sex , age")
		fmt.Println("举个例子：    Leo,0001,man,12 , 必须以逗号分隔")

		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happed a error:", err)
			return
		}
		rawData := strings.Split(string(result), ",")
		name := strings.TrimSpace(rawData[0])
		if name == "quit" {
			fmt.Println("exit to the add student")
			break
		}
		if len(rawData) != 4 {
			fmt.Println("you weren't provide some right arguments")
			continue
		}
		ID, _ := strconv.Atoi(strings.TrimSpace(rawData[1]))
		sex := strings.TrimSpace(rawData[2])
		age, _ := strconv.Atoi(strings.TrimSpace(rawData[3]))
		if name == "" && ID == 0 && sex == "" && age == 0 {
			continue
		}
		stu := model.Student{
			Name: name,
			ID:   ID,
			Sex:  sex,
			Age:  age,
		}
		stuJson[i] = stu
		i++
	}
	model.SaveStuData(stuJson)
}

func BorrowBook() {
	// 借书的功能

}

func main() {

	msg := `
1. 添加图书信息
2. 添加学生信息
3. 借书
4. 后台管理（删除书籍与学生信息）
5. 显示当前注册的图书
6. 显示当前注册的学生信息`
	for {
		fmt.Println(msg)
		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error: ", err)
			return
		}
		input := strings.TrimSpace(string(result))
		if len(input) != 1 {
			continue
		}
		i, _ := strconv.Atoi(string(input[0]))

		switch i {
		case 1:
			AddBook()
		case 2:
			AddStu()

		case 3:
			BorrowBook()
			//		case 4:
			//			borrowBook(&bInit, &stuInit)

			//		case 5:
			//			management(&bInit, &stuInit)
			//		case 6:
			//			showBook(&bInit)
			//		case 7:
			//			showStu(&stuInit)
		default:
			fmt.Println("you weren't input a available choice!!")
			continue
		}
	}
}
