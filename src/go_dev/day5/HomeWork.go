// HomeWork
package main

import (
	"fmt"
	"strconv"
	"strings"
	//	"strings"
	"bufio"
	"os"
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
	for b != nil {
		fmt.Println("book name", b.name, ", book author", b.author, ", book published",
			b.published, ", how many book", b.many, ", the next book", b.next)

		b = b.next
		time.Sleep(100 * time.Millisecond)
	}
}

func addBook(b *book) book {
	// 增加书籍的，录入书籍的信息的格式必须是  书名,作者,出版时间,多少本   必须以逗号分隔。

	var bb = b

	for {
		fmt.Println("please input the book's basic info,enter quit will quit this func, and the format is ")
		fmt.Println(" name, author ,published , how many")
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
		bb.next = &bookInfo
		bb = &bookInfo

	}
	return *b
}

func selectBook(b *book, key string) {
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
			fmt.Println("the book name", b.name, ", the book author", b.author, ", the book published",
				b.published, ",and how many book", b.many)
		}
		b = b.next
	}
	//	}
}

func addStudent(s *student) *student {
	// 录入学生信息的
	var st = s
	var name, ID, sex string

	for {
		reader := bufio.NewReader(os.Stdin)
		rt, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("happend a error: ", err)
			continue
		}
		rt = strings.Split(string(rt), ",")
		name = strings.TrimSpace(rt[0])
		if name == "quit" {
			fmt.Println("exit to add student's info")
			break
		}
		if len(rt) != 3 {
			fmt.Println("you weren't provide right arguments")
			continue
		}
		ID = strings.TrimSpace(rt[1])
		sex = strings.TrimSpace(rt[2])
		stu := student{
			name: name,
			ID:   ID,
			sex:  sex,
		}
		st.next = &stu
		st = &stu
	}
	return s
}

func main() {
	bInit := book{
		name:      "a journey to west",
		author:    "Mr.Wu",
		published: "1986-07-01",
		many:      12,
	}
	for {
		msg := `
1： add books
2: select book
3: add students's info
4: borrow books
5: manage books
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
			bInit = addBook(&bInit)
			showBook(&bInit)

		case 2:
			fmt.Println("select a book from library,you can input the book's name,author or published! enter quit will exit")
			rt := bufio.NewReader(os.Stdin)
			input, _, _ := rt.ReadLine()
			key := string(input)
			fmt.Println("keyaa", key)
			if key == "quit" {
				return
			}
			selectBook(&bInit, key)

		case 3:

		default:
			fmt.Println("you weren't input a available choice!!")
			continue
		}
	}
}
