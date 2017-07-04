// HomeWork
package main

import (
	"fmt"
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
	for b != nil {
		fmt.Println("book name", b.name, "book author", b.author, "book published",
			b.published, "how many book", b.many, "the next book", b.next)

		b = b.next
		time.Sleep(100 * time.Millisecond)
	}
}

func addBook(b *book) book {
	var bb = b

	var (
		name, author, published string
		many                    int
	)

	for {
		fmt.Println("please input the book's basic info,q is quit, and the format is ")
		fmt.Println(" name, author ,published , many")
		fmt.Scanf("%s,%s,%s,%d", &name, &author, &published, &many)

		if strings.HasPrefix(name, "q") && author == "" && published == "" && many == 0 {
			fmt.Println("exit to the add Book")
			break
		}
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
		showBook(b)
	}
	return *b
}

func main() {
	bInit := book{
		name:      "a journey to west",
		author:    "Mr.Wu",
		published: "1986-07-01",
		many:      12,
	}

	addBook(&bInit)
}
