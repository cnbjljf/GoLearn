package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// flag.String用来接收命令行参数的，这个方法第一个参数是短写的命令，第二个参数是接收到命令后赋值给这个变量
// 第三个参数是这命令的解释
var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to recevie sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	// 这个方法用来读取输入的文件内容，并重组成int型数组返回
	file, err := os.Open(infile) // 打开文件
	if err != nil {              // 如果打开文件没有错误的话
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
	if infile != nil { // 如果infile 不为空，那么就执行下面的代码
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithms =>", *algorithm)
	}
	value, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read value", value)
	} else {
		fmt.Println(err)
	}
}
