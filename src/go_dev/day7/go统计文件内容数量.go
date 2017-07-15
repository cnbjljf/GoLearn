// go统计文件内容数量
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("aaa.go")
	reader := bufio.NewReader(f)
	var azNum int
	var intNum int
	var fuhaoNum int
	var otherNum int
	for {
		rawstring, err := reader.ReadString('\n')
		if err != io.EOF {
			break
		}
		result := []rune(rawstring)
		for _, key := range result {
			//fmt.Printf("%v %T\n", key, key)
			switch {
			case key >= 'a' && key <= 'z':
				fallthrough
			case key >= 'A' && key <= 'Z':
				azNum++
			case key >= '0' && key <= '9':
				intNum++
			case key == ' ' || key == '\t':
				fuhaoNum++
			default:
				otherNum++
			}
		}
	}
	fmt.Println("azNum", azNum)
	fmt.Println("intNum", intNum)
	fmt.Println("fuhaoNum", fuhaoNum)
	fmt.Println("otherNum", otherNum)
}
