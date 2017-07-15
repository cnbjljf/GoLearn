// WriteToFile
package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	outputFile,err := os.OpenFile("output.txt",os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		fmt.Printf("an error happend with fil crea\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i:=0;i<10;i++ {
		outputWriter.WriteString(outputString)
	}
	outputFile.
}
