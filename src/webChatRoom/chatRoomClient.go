// chatRoomClient
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func messageSend(conn net.Conn) {
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data, _, err := reader.ReadLine()
		CheckError(err)
		input = string(data)
		if strings.ToLower(input) == "exit" {
			conn.Close()
			break
		}
		_, err = conn.Write(data)
		if err != nil {
			conn.Close()
			fmt.Println("Client connect failure: ", err.Error())
			break
		}
	}
	defer conn.Close()
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer conn.Close()

	go messageSend(conn)
	//	conn.Write([]byte("hello,i am Leo\n"))
	for {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if length != 0 && err == nil {
			fmt.Println("receivce a message ", strings.TrimSpace(string(buf)))
		} else {
			continue
		}
	}

	fmt.Println("exit the program!!!")
}
