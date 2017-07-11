// main
package main

import (
	"io/ioutil"
	"strings"
	"bufio"
	"encoding/json"
	"fmt"
	"go_dev/day6/HomeWork/model/book"
	"os"
)

const (
	StockNum 10000   // 图书馆图书容量
	bookDataFilePath "d:/bookData.json"
	stuDataFilePath "d:/studentData.json"
)

var bookJson map[int]book.Book

func savebookJsonData(bk bookJson) bool{
	// 保存图书信息到文本文件的
	// 先读取文件里买的数据，然后在写入，避免覆盖之前的数据

	
	
	f , err := os.Open(bookDataFilePath)
	
	if err != nil {
		fmt.Println("happend a error:",err)
		return false
	}
	bkJsonData := json.Marshal(bookJson)
	
}



func AddBook(){
	bookJson = make(map[int]book.Book)
	// 增加书籍的，录入书籍的信息的格式必须是  书名,作者,出版时间,多少本   必须以逗号分隔。
	oldF , err := ioutil.ReadFile(bookDataFilePath)	
	if err != nil {
		fmt.Println("happend a error:",err)
		return false
	}
	oldData , _ := json.Unmarshal(oldF,&bookJson)


	var i = 0
	for i<StockNum {
		fmt.Println("Please Input your book's infoenter quit will quit this func, and the format is ")
		fmt.Println(" name, author ,published , how many")
		fmt.Println("举个例子：    aaa,Leo,20170705,12")
		
		reader := bufio.NewReader(os.Stdin)
		result,_,err := reader.ReadLine()
		if err != nil {
			fmt.Println("happed a error:",err)
			return 
		}
		rawData := strings.Split(string(result),",")
		bookName = strings.TrimSpace(rawData[0])
		if name == "quit" {
			fmt.Println("exit to the add Book")
			break
		}
		if len(raw_data) != 4 {
			fmt.Println("you weren't provide some right arguments")
			continue
		}
		author := strings.TrimSpace(raw_data[1])
		published := strings.TrimSpace(raw_data[2])
		many, _ := strconv.Atoi(strings.TrimSpace(raw_data[3]))
		if name == "" && author == "" && published == "" && many == 0 {
			continue
		}
		bookInfo := book.Book{
			Name       name
			Total      many
			Author     author
			CreateTime published
		}
//		bookJsonInfo,err = json.Marshal(bookInfo)
//		if  err!=nil {
//			fmt.Println("happed a error:",err)
//			return 
//		}
		oldData[]
		
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
	fmt.Println(msg)
		if err != nil {
			fmt.Println("happend a error: ", err)
			return
		}
		if len(input) != 1 {
			continue
		}
		i, _ := strconv.Atoi(string(input[0]))

		switch i {
		case 1:
			

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
