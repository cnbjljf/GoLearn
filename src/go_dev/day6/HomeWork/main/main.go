// main
package main

import (
	"bufio"
	"fmt"
	"go_dev/day6/HomeWork/constValue" // 定义一些常量，相当于配置文件一样
	"go_dev/day6/HomeWork/model"
	"log"
	"os"
	"sort"
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

	bookJson, err := model.GetBookOldData() //先判断文件是否存在，存在才读取老文件
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

func BorrowBook(userName string, stuData map[int]model.Student) {
	// 借书的功能
	var choice string
	for {
		showDetail("book")

		fmt.Println("请输入书名,每次只能借一本")
		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		choice = strings.TrimSpace(string(result))
		data, err := model.GetBookOldData()
		if err != nil {
			fmt.Println("show the detail of book's info happend a error:", err)
		}

		if choice == "quit" {
			break
		}
		for i, v := range data { // 遍历图书map
			if choice == v.Name {
				if v.Total-1 >= 0 { // 输入正确的图书名且图书存量大于1，那么可以外借
					for ii, tmp := range stuData { // 找到这个学生的信息，
						if userName == tmp.Name {
							//把借书的信息和学生信息进行绑定
							stuV := stuData[ii]
							stuV.BrrowBook = append(stuV.BrrowBook, v)
							stuData[ii] = stuV
							fmt.Println(stuData)
							model.SaveStuData(stuData)
						}
					}
					// 下面几步对图书存量-1
					tmp := data[i]
					tmp.Total--
					data[i] = tmp
					//					fmt.Println(data[i])
					model.SavebookJsonData(data)
					fmt.Println("借书成功，书名:[", choice, "]该本书存量:[", data[i].Total, "]")
				} else {
					fmt.Println("该本书没有剩余了，请等待他人归还再借！")
				}
			}
		}
	}
}

func ReturnBook(name string, stuData map[int]model.Student) {
	// 还书的功能
	fmt.Printf("这些是你借到的书籍：\n")
	var BorrowBookNum map[string]int // 统计每本书借了多少本
	BorrowBookNum = make(map[string]int)
	for {
		for _, v := range stuData {
			if name == v.Name {
				for _, book := range v.BrrowBook {
					value, ok := BorrowBookNum[book.Name]
					if ok { // 意味着有值
						BorrowBookNum[book.Name] = value + 1
					} else {
						BorrowBookNum[book.Name] = 1
					}
				}
				fmt.Printf("%-13s%-5s", "书名", "数量\n")
				fmt.Println()
				for k, v := range BorrowBookNum {
					fmt.Printf("%-15s%-5d\n", k, v)
				}
			}
		}
		fmt.Printf("\n请输入书名\n")
		reader := bufio.NewReader(os.Stdin)
		result, _, _ := reader.ReadLine()
		bookName := strings.TrimSpace(string(result))
		if bookName == "quit" {
			return
		}
		_, ok := BorrowBookNum[bookName]
		if ok {
			for k, v := range stuData {
				if name == v.Name {
					var leftBook []model.Book
					for _, book := range v.BrrowBook {
						if bookName != book.Name {
							leftBook = append(leftBook, book)
						}
					}
					fmt.Println(leftBook)
					fmt.Println(v.BrrowBook)
					v.BrrowBook = leftBook
					stuData[k] = v
					fmt.Println("还书成功！！！")
				}
			}
		} else {
			fmt.Println("没有找到这本书在你已借书列表里面，请确认书名！！")
		}
	}

	return
}

func showDetail(name string) {
	// 根据name值来显示图书或者学生的详细信息
	var idList []int
	if name == "student" {
		data, err := model.GetStuOldData()
		if err != nil {
			fmt.Println("show the detail of student's info happend a error:", err)
		}

		for id, _ := range data {
			idList = append(idList, id)
		}
		sort.Ints(idList)
		fmt.Println("======================= student ===================================")
		for _, k := range idList {
			item := data[k]
			fmt.Printf("%d. name: %-8s ,id: %-5d,sex: %-10s,age: %-2d\n", k, item.Name, item.ID, item.Sex, item.Age)
		}
		fmt.Println("==================================================================")
	} else if name == "book" {
		data, err := model.GetBookOldData()
		if err != nil {
			fmt.Println("show the detail of book's info happend a error:", err)
		}
		for id, _ := range data {
			idList = append(idList, id)
		}
		sort.Ints(idList)
		fmt.Println("====================== BOOK ======================================")
		for _, k := range idList {
			item := data[k]
			fmt.Printf("%d. name: %-8s ,author: %-8s ,stock: %-5d,published time: %-12s\n",
				k, item.Name, item.Author, item.Total, item.CreateTime)
		}
		fmt.Println("==================================================================")
	}
}

func manage() {
	// 管理图书/学生信息的，主要是删除或者修改
}

func main() {
	var name string
	var pwd string
	fmt.Println("请输入用户名")
	fmt.Scanln(&name)
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	_, ok := constValue.LoginAdmin(name, pwd)
	if ok {
		for {
			fmt.Println(constValue.AdminMsg)
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
				manage()
			case 4:
				showDetail("student")
			case 5:
				showDetail("book")
			case 6:
				return
			default:
				fmt.Println("you weren't input a available choice!!")
				continue
			}
		}
	}
	pwdID, _ := strconv.Atoi(pwd)
	data, ok := model.LoginStu(name, pwdID)
	if ok {
		for {
			fmt.Println(constValue.CustonmerMsg)
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
				BorrowBook(name, data)
			case 2:
				ReturnBook(name, data)
			case 3:
				return
			default:
				fmt.Println("you weren't input a available choice!!")
				continue
			}
		}
	}
	fmt.Println("错误的用户名或者密码！！！")
}
