// HomeWork
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type book struct {
	name      string
	author    string
	published string
	many      int
	next      *book
}

type student struct {
	name  string
	grade int
	ID    string
	sex   string
	book  *book
	next  *student
}

func showBook(b *book) {
	// 显示书籍的
	for b != nil {
		fmt.Println("book name", b.name, ", book author", b.author, ", book published",
			b.published, ", how many book", b.many, ", the next book", b.next)
		b = b.next
		time.Sleep(10 * time.Millisecond)
	}
}

func showStu(s *student) {
	// 显示学生信息的
	for s != nil {
		fmt.Println("student's name", s.name, ", student's sex", s.sex, ", student's ID",
			s.ID, "the borrowed book is [", s.book, "] the next student", s.next)

		s = s.next
		time.Sleep(10 * time.Millisecond)
	}
}

func addBook(b *book) *book {
	// 增加书籍的，录入书籍的信息的格式必须是  书名,作者,出版时间,多少本   必须以逗号分隔。

	for {
		var bbb = b
		fmt.Println("please input the book's basic info,enter quit will quit this func, and the format is ")
		fmt.Println(" name, author ,published , how many")
		fmt.Println("举个例子：    aaa,Leo,20170705,12")
		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error!", err)
			break
		}
		raw_data := strings.Split(string(result), ",")

		name := strings.TrimSpace(raw_data[0])
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
		bookInfo := book{
			name:      name,
			author:    author,
			published: published,
			many:      many,
		}
		previous := bbb
		for bbb != nil {
			previous = bbb
			bbb = bbb.next
		}

		previous.next = &bookInfo

	}
	return b
}

func selectBook(b *book, key string) (*book, bool) {
	// 查询书籍的

	//	fmt.Println(name, author, published)
	//	for {
	//		fmt.Println("select a book from library,you can input the book's name,author or published! enter quit will exit")
	//		rt := bufio.NewReader(os.Stdin)
	//		input, _, _ := rt.ReadLine()
	//		key := string(input[0])
	//		if key == "quit" {
	//			return
	//		}

	for b != nil {
		if b.name == key || b.author == key || b.published == key {
			fmt.Println("find this book ,the detail of the book will show you ")
			fmt.Println("the book name is [", b.name, "], the book author is[", b.author, "], the book published is[",
				b.published, "],and we have [", b.many, "] book")
			return b, true
			break
		}
		b = b.next
	}
	return b, false

	//	}
}

func addStudent(s *student) *student {
	// 录入学生信息的

	var name, ID, sex string

	for {
		var st = s
		fmt.Println("please input student's info,enter quit is exit,and the format is name,ID,sex    ")
		fmt.Println("必须以逗号分隔！！！例如：   ljf,2017070501,man")
		reader := bufio.NewReader(os.Stdin)
		rt, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error: ", err)
			continue
		}
		result := strings.Split(string(rt), ",")
		name = strings.TrimSpace(result[0])
		if name == "quit" {
			fmt.Println("exit to add student's info")
			break
		}
		if len(result) != 3 {
			fmt.Println("you weren't provide right arguments")
			continue
		}
		ID = strings.TrimSpace(result[1])
		sex = strings.TrimSpace(result[2])
		stu := student{
			name: name,
			ID:   ID,
			sex:  sex,
		}
		previous := st
		for st != nil {
			previous = st
			st = st.next
		}
		previous.next = &stu

	}
	return s
}

func borrowBook(b *book, s *student) student {
	// 借书的功能，主要修改学生的信息
	var bb = b
	var ss = s

	// 先选书籍
	fmt.Printf("\n\nthese books are left in library!\n")
	for {
		var i int
		for b != nil {
			i++
			fmt.Printf("%d. %s[author:%s,left:%d]\n", i, b.name, b.author, b.many)
			b = b.next
		}

		fmt.Println("please input which book do you want to borrow, you can type  book's name down!enter quit will be quit")
		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error!\n", err)
		}
		bookName := strings.TrimSpace(string(result))

		if bookName == "quit" {
			break
		}

		// 再选哪个学生借书了
		fmt.Printf("\n\nthese student's were enroll!!!\n\n")
		i = 0
		for s != nil {
			i++
			fmt.Printf("%d. %s(ID:%s)\n", i, s.name, s.ID)
			s = s.next
		}

		fmt.Println("please input the student's name to make sure who borrow the book!enter quit will be quit")
		reader = bufio.NewReader(os.Stdin)
		result, _, err = reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error!", err)
		}
		stuName := strings.TrimSpace(string(result))

		if strings.ToLower(stuName) == "quit" {
			break
		}

		// flag为标志位，判断是否找到了指定名字的书籍
		brBook, flag := selectBook(bb, bookName)
		if flag {
			for ss != nil {
				if ss.name == stuName {
					brBook.many--
					if brBook.many >= 0 {
						ss.book = brBook
						fmt.Printf("the student[%s] borrow the book[%v]\n", ss.name, *ss.book)
						break
					} else {
						fmt.Println("sorry, the library hasn't enought book to borrow you!")
						break
					}
				}
				ss = ss.next
			}
			break
		} else {
			fmt.Println("not found which book is you look for")
			break
		}
	}

	return *ss
}

func delStuNode(s *student) {
	// 删指定学生
	var ss = s
	fmt.Println("these students are enrolled!")
	for s != nil {
		fmt.Printf("the student's name:[%s] id[%s]\n", s.name, s.ID)
		s = s.next
	}
	fmt.Println("please input the student's name to delete ")
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("happend a error", err)
		return
	}
	chooseName := strings.TrimSpace(string(result))
	if strings.ToLower(chooseName) == "quit" {
		return
	}
	previous := ss
	for ss != nil {
		if ss.name == chooseName {
			previous.next = ss.next
		}
		previous = ss
		ss = ss.next
	}
}

func delBkNode(b *book) {
	// 删指定书
	var bb = b
	fmt.Printf("\nthese books are enrolled!!\n")
	for b != nil {
		fmt.Printf("the book name: [%s] author:[%s]\n", b.name, b.author)
		b = b.next
	}
	fmt.Printf("\nplease input the books name!\n")
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("happend a error!", err)
	}
	inputName := strings.TrimSpace(string(result))
	if strings.ToLower(inputName) == "quit" {
		return
	}
	previous := bb
	for bb != nil {
		if bb.name == inputName {
			previous.next = bb.next
		}
		previous = bb
		bb = bb.next
	}

}

func management(b *book, s *student) {
	// 后台管理，删除学生信息与书本信息的
	for {
		fmt.Println("Please choose the serial num,quit will be quit!")
		msg := `
	1. manage book
	2. manage student
	`
		fmt.Println(msg)

		reader := bufio.NewReader(os.Stdin)
		result, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error", err)
		}
		rt := strings.TrimSpace(string(result))
		if strings.ToLower(rt) == "quit" {
			break
		}
		inputNum, _ := strconv.Atoi(rt)
		switch inputNum {
		case 1:
			delBkNode(b)
		case 2:
			delStuNode(s)
		}
	}
}

func main() {
	//  初始化管理员信息与第一本书籍信息
	bInit := book{
		name:      "a journey to west",
		author:    "Mr.Wu",
		published: "1986-07-01",
		many:      1,
	}

	stuInit := student{
		name: "admin",
		sex:  "man",
		ID:   "2017070501",
	}

	for {
		msg := `
		1：add books
		2: select book
		3: add students's info
		4: borrow books
		5: manage books(del a book or a student!)
		6: show books
		7: show students
		`
		fmt.Println(msg)
		reader := bufio.NewReader(os.Stdin)
		input, _, err := reader.ReadLine()
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
			showBook(addBook(&bInit))

		case 2:
			fmt.Println("select a book from library,you can input the book's name,author or published! enter quit will exit")
			rt := bufio.NewReader(os.Stdin)
			input, _, _ := rt.ReadLine()
			key := string(input)
			if strings.ToLower(key) == "quit" {
				break
			}
			selectBook(&bInit, key)

		case 3:
			showStu(addStudent(&stuInit))
		case 4:
			borrowBook(&bInit, &stuInit)

		case 5:
			management(&bInit, &stuInit)
		case 6:
			showBook(&bInit)
		case 7:
			showStu(&stuInit)

		default:
			fmt.Println("you weren't input a available choice!!")
			continue
		}
	}
}
