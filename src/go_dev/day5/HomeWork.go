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

func selectBook(b *book) {
	// 查询书籍的
	var name, author, published string
	//	fmt.Println(name, author, published)
	for {
		fmt.Println("select a book from library,you can input the book's name,author or published! enter quit will exit")
		rt := bufio.NewReader(os.Stdin)
		input, _, _ := rt.ReadLine()
		key := string(input[0])
		if key == "quit" {
			return
		}
		for b != nil {
			if b.name == key || b.author == key || b.published == key {
				fmt.Println(", book name", b.name, ", book author", b.author, ", book published",
					b.published, ", how many book", b.many, ", the next book", b.next)
			}
			b = b.next
		}
	}
}

func main() {
	bInit := book{
		name:      "a journey to west",
		author:    "Mr.Wu",
		published: "1986-07-01",
		many:      12,
	}
	var bk book
	fmt.Println(bInit, bk)
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
		i, _ := strconv.Atoi(string(input[0]))
		fmt.Printf("%T,%v", i, i)
		switch i {
		case 1:
			bk = addBook(&bInit)
			showBook(&bk)
		case 2:
			selectBook(&bk)
		default:
			fmt.Println("you weren't input a available choice!!")
			continue
		}
	}
}
