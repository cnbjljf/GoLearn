// main
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	//	"go_dev/day6/HomeWork/model/book"
	"os"
)

func main() {

	msg := `
1. 添加图书信息
2. 添加学生信息
3. 借书
4. 后台管理（删除书籍与学生信息）
5. 显示当前注册的图书
6. 显示当前注册的学生信息`
	fmt.Println(msg)
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	fmt.Println(result)

	f, err := os.Create("d:/student_info.txt")
	if err != nil {
		fmt.Println("fail to open file student_info.txt ")
	}
	defer f.Close()
	var he map[int]string
	he = make(map[int]string)
	he[1] = "1111"
	he[2] = "2222"

	hehe, _ := json.Marshal(he)

	f.Write(hehe)
	f.Sync()

}
