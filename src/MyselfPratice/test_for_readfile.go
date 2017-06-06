package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil { // 如果打开文件没有错误的话
		fmt.Println("Faild to open the input file...", infile)
		return
	}
	defer file.Close()          // 如果文件操作失败，那么无论如何都要关闭文件句柄
	br := bufio.NewReader(file) // 读取这个文件
	values = make([]int, 0)     // 定义一个数组

	for {
		line, isPrefix, err1 := br.ReadLine() // 读取每一行
		fmt.Println(line, isPrefix, err1)
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line,seems unexpected")
			return
		}
		str := string(line) // 转换字符数组为字符串
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func main() {
	flag.Parse()
	value, err := readValues(*infile)
	fmt.Println("values:", value, "error--->", err)

}
