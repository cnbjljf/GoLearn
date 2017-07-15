// copyFile
package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("happend a error when opening", err)
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("happend a error when opening", err)
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func main() {
	copyFile("test.gz", "test.log.gz")
}
