//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {

	var aa map[int]Book
	aa = make(map[int]Book)
	//	aa[1] = "aaa"
	//	aa[2] = "bbb"
	//	d, _ := json.Marshal(aa)
	//	//		fmt.Println(aa)
	//	f, _ := os.Create("d:\test.json")
	//	fmt.Println(d)
	//	f.Write(d)
	//	f.Sync()
	//	f.Close()

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

	//	ab, _ := ioutil.ReadFile("d:/student_info.txt")
	//	json.Unmarshal(ab, &aa)
	//	fmt.Println(aa)
	//	aa[3] = "cc"
	//	fmt.Println(aa)
	//	fmt.Println(aa[3])

	//	var i = 0
	//	for i < 100 {
	//		_, ok := aa[i]
	//		fmt.Println(i, ok)
	//		i++
	//	}
	//	fmt.Println(i + 1)
	//	for i < 200 {
	//		fmt.Println(i)
	//		i++
	//	}
	filename := "d:/bookData.json"
	fmt.Println(Exist(filename))
	fmt.Println(os.Stat(filename))
	fmt.Println()
	_, err := os.Stat(filename)
	fmt.Println(os.IsExist(err), os.IsNotExist(err))
	if err != nil || os.IsNotExist(err) {
		fmt.Println("file  exist")
		ab, _ := ioutil.ReadFile(filename)
		json.Unmarshal(ab, &aa)
		//		fmt.Println(aa)
		for k, v := range aa {
			fmt.Println(k, v)
		}
		//		aa[3] = "cc"
		//		fmt.Println(aa)
		//		fmt.Println(aa[3])

	}
	fmt.Println("done")
	//		filename := "a-nonexistent-file"
	//		if _, err := os.Stat(filename); os.IsNotExist(err) {
	//		    fmt.Printf("file does not exist")
	//		}

}
