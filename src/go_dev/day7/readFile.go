package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	inputFile, err := os.Open("aaa.go")
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}
