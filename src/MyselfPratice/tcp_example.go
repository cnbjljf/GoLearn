// tcp_example
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	fmt.Println("开始建立连接")
	conn, err := net.Dial("tcp", service)
	checkError(err)
	fmt.Println("开始发送head信息")
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Printf("%s 返回来的信息如下: \n\n", service)
	fmt.Println(string(result))

	os.Exit(1)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}

}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
