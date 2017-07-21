//
package main

import (
	"time"
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"os"
)

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

type stu struct {
	Name  string
	bbook []book
}

type book struct {
	Name string
}

//func main() {
//	var bl []book
//	for i := 0; i < 6; i++ {
//		var b = book{
//			Name: fmt.Sprintf("stu%d", i),
//		}
//		bl = append(bl, b)
//	}

//	var ss = stu{
//		Name:  "first",
//		bbook: bl,
//	}
//	fmt.Println(ss)
//	var b2 []book
//	for i := 10; i < 16; i++ {
//		var b = book{
//			Name: fmt.Sprintf("stu%d", i),
//		}
//		b2 = append(b2, b)
//	}
//	ss.bbook = b2
//	fmt.Println(ss)
//}

func Logger(username, action string, content interface{}) {
	ct, ok := content.(string)
	if ok == false {
		fmt.Println(ok)
		return
	}
	const LogFile = "d:/info.log"
	f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("opening file happend a error!,", err)
		return
	}
	logContent := fmt.Sprintf("%s user:[%s] action:[%s] info[%s]\n",
		time.Now().Format("2006-01-02 15:04:05"), username, action, ct)
	f.WriteString(logContent)
	fmt.Println(logContent)
	defer f.Close()
}

func main() {

	var ss = book{
		Name: "first",
	}
	Logger("admin", "delete", ss)
}
