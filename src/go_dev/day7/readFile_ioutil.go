// readFile_ioutil
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "d:/111.txt"
	outputFile := "d:/112.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error:%s", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
