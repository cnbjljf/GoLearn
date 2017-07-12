// main
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_dev/day6/HomeWork/model"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	StockNum         = 10000 // 图书馆图书容量
	bookDataFilePath = "d:/bookData.json"
	stuDataFilePath  = "d:/studentData.json"
)

var bookJson map[int]model.Book

func Exist(filename string) bool { // 判断指定文件是否存在的
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func savebookJsonData(bk map[int]model.Book) bool {
	// 保存图书信息到文本文件的
	f, err := os.Create(bookDataFilePath)

	if err != nil {
		log.Fatalln("Saving data  happend a error:", err)
		return false
	}
	bkJsonData, _ := json.Marshal(bookJson)
	f.Write(bkJsonData)
	f.Sync()
	defer f.Close()
	fmt.Println("Saving data successfully!! these book infomation is :")
	for k, v := range bk {
		fmt.Printf("ID: %d, book:%v\n", k, v)
	}
	return true
}

func AddBook() {
	// 增加书籍的，录入书籍的信息的格式必须是  书名,作者,出版时间,多少本   必须以逗号分隔。
	// 先读取文件里买的数据，然后在写入，避免覆盖之前的数据
	bookJson = make(map[int]model.Book)

	if Exist(bookDataFilePath) { // 先判断文件是否存在，存在才读取老文件
		oldF, err := ioutil.ReadFile(bookDataFilePath) // 读取文本数据
		if err != nil {
			log.Fatalln("happend a error:", err)
			return
		}
		fmt.Println("find the  book's  data file and then ready to load it!!these are exist book!")
		json.Unmarshal(oldF, &bookJson) // 加载之前的数据
		for k, v := range bookJson {
			fmt.Printf("ID: %d, book:%v\n", k, v)
		}
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
		//		bookJsonInfo,err = json.Marshal(bookInfo)
		//		if  err!=nil {
		//			fmt.Println("happed a error:",err)
		//			return
		//		}
		bookJson[i] = bookInfo
		i++
	}
	if savebookJsonData(bookJson) {
		return
	} else {
		fmt.Println()
	}

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
			//		case 2:

			//		case 3:
			//			showStu(addStudent(&stuInit))
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

		//	f, err := os.Create("d:/student_info.txt")
		//	if err != nil {
		//		fmt.Println("fail to open file student_info.txt ")
		//	}
		//	defer f.Close()
		//	var he map[int]string
		//	he = make(map[int]string)
		//	he[1] = "1111"
		//	he[2] = "2222"

		//	hehe, _ := json.Marshal(he)

		//	f.Write(hehe)
		//	f.Sync()
	}
}
