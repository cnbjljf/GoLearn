// main
package main

import (
	"bufio"
	"fmt"
	"go_dev/day6/HomeWork/constValue" // 定义一些常量，相当于配置文件一样
	"go_dev/day6/HomeWork/model"
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

func AddBook(username string) {

	// 增加书籍的，录入书籍的信息的格式必须是  书名,作者,出版时间,多少本   必须以逗号分隔。
	// 先读取文件里买的数据，然后在写入，避免覆盖之前的数据
	bookJson = make(map[int]model.Book)

	bookJson, err := model.GetBookOldData() //先判断文件是否存在，存在才读取老文件
	if err != nil {
		fmt.Println(err)
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
		fmt.Println("举个例子：    aaa,Leo,20170705,12 , 必须以逗号分隔，输入quit退出")

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
		constValue.Logger(username, "add a book", bookInfo)
		bookJson[i] = bookInfo
		i++
	}
	if model.SavebookJsonData(bookJson) {
		return
	} else {
		fmt.Println("Save data happend a error!")
	}
}

func AddStu(username string) {
	// 添加学生信息的，录入学生的信息必须是i
	stuJson = make(map[int]model.Student)

	stuJson, err := model.GetStuOldData() //先判断文件是否存在，存在才读取老文件
	if err != nil {
		fmt.Println(err)
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
		fmt.Println("举个例子：    Leo,0001,man,12 , 必须以逗号分隔，输入quit退出")

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
		constValue.Logger(username, "add a student", stu)
		i++
	}

	model.SaveStuData(stuJson)
}

func BorrowBook(userName string, stuData map[int]model.Student) {
	// 借书的功能
	/*
		userName 谁登陆的

	*/
	var choice string
	for {
		showDetail("book")

		fmt.Println("请输入书名,每次只能借一本，输入quit退出")
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
					constValue.Logger(userName, "borrow a book", "the book name is "+choice)
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

	for {
		BorrowBookNum = make(map[string]int)
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
		fmt.Printf("\n请输入书名,每次只能还一本输入quit退出\n")
		reader := bufio.NewReader(os.Stdin)
		result, _, _ := reader.ReadLine()
		bookName := strings.TrimSpace(string(result))
		if bookName == "quit" {
			return
		}
		returnBookNum := 1                 // 保留字段，用来调整每次默认还多少本书的数量，默认为1
		num, ok := BorrowBookNum[bookName] // value 是还有多少本没有还
		if ok {
			for k, v := range stuData {
				if name == v.Name {
					var leftBook []model.Book // 用来存放还有哪些书没有归还
					var i int
					for _, book := range v.BrrowBook {
						// 如果书名相等并且书籍数量小于归还后剩余的数量，否则添加到leftBook数组里
						if bookName == book.Name && i < num-returnBookNum {
							leftBook = append(leftBook, book)
							i++
						} else if bookName != book.Name {
							leftBook = append(leftBook, book)
						}

					}
					v.BrrowBook = leftBook
					stuData[k] = v
					model.SaveStuData(stuData)
					bookData, _ := model.GetBookOldData()
					for k, v := range bookData {
						if v.Name == bookName { // 找到这本书
							v.Total = v.Total + returnBookNum
							bookData[k] = v
						}
					}
					constValue.Logger(name, "return a book", "book name is "+bookName)
					model.SavebookJsonData(bookData)
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

func manage(username string) {
	// 管理图书/学生信息的，主要是删除或者修改
	fmt.Println(constValue.ManageChoice)
	inputer := bufio.NewReader(os.Stdin)
	result, _, _ := inputer.ReadLine()
	inputNum, _ := strconv.Atoi(strings.TrimSpace(string(result)))
	switch inputNum {
	case 1:
		model.ManageStudent(username)
	case 2:
		model.ManageBook(username)
	}
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
				constValue.Logger(name, "add a book", "begin to add a book")
				AddBook(name)
			case 2:
				constValue.Logger(name, "add a student", "begin to add a Student")
				AddStu(name)
			case 3:
				constValue.Logger(name, "manage ", "going to manage platform")
				manage(name)
			case 4:
				constValue.Logger(name, "show", "show the student's info")
				showDetail("student")
			case 5:
				constValue.Logger(name, "show", "show the book's info")
				showDetail("book")
			case 6:
				constValue.Logger(name, "exit", "exit the program!!!")
				return
			default:
				constValue.Logger(name, "unavailable", "input a unavailable choice!!")
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
				constValue.Logger(name, "borrow a book", "begin to borrow a book!!!")
				BorrowBook(name, data)
			case 2:
				constValue.Logger(name, "return a book", "begin to return a book!!!")
				ReturnBook(name, data)
			case 3:
				constValue.Logger(name, "exit", "exit the program!!!")
				return
			default:
				constValue.Logger(name, "unavailable", "input a unavailable choice!!")
				fmt.Println("you weren't input a available choice!!")
				continue
			}
		}
	}
	constValue.Logger(name, "login", "unavailable username or password!!")
	fmt.Println("错误的用户名或者密码！！！")
}
