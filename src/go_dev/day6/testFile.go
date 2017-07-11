//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//	"os"
)

func main() {
	var aa map[int]string
	aa = make(map[int]string)
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

	ab, _ := ioutil.ReadFile("d:/student_info.txt")
	json.Unmarshal(ab, &aa)
	fmt.Println(aa)
	aa[3] = "cc"
	fmt.Println(aa)

}
