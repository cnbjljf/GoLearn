// gzipFile
package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fName := "test.log.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v,Can't open %s: \n")
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		y
		fmt.Fprintf(os.Stderr, "open zip file faild,err:%v\n", err)
		return
	}
	r = bufio.NewReader(fz)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
