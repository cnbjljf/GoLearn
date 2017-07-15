// Readstring
package main

import (
	"bufio"
	"fmt"
	"os"
)

const name = "E:\111.txt"

func openFile(name string) {
	f, _ := os.Open(name)
	newf := bufio.NewReader(f)
	result, _, err := newf.ReadLine()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(result)
}

func main() {
	//	str := bufio.NewReader(os.Stdin)
	//	resultt, err := str.ReadString('\n')
	//	if err != nil {
	//		fmt.Println("read string failed ,err:", err)
	//		return
	//	}
	//	fmt.Printf("read the str successfully,ret:%s\n", resultt)
	openFile(name)
}
