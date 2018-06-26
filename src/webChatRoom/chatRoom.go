// chatRoom
package main

import (
	"fmt"
	"net"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

var onlineConns = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quiteChan = make(chan string, 1000)

func ProcessInfo(conn net.Conn) {

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		if numOfBytes != 0 {
			remoteAddr := conn.RemoteAddr()
			fmt.Printf("has received this message:%s:%s\n", remoteAddr, strings.TrimSpace(string(buf)))
			messageQueue <- string(buf)

		}
	}
}

func ConsumMessage() {
	for {
		select {
		case message := <-messageQueue:
			// 对消息进行解析
			doProcessMessage(message)

		case <-quiteChan:
			break
		}
	}
}

func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		addr := contents[0]
		sendMessages := strings.TrimSpace(contents[1 : len(contents)-1])
		addr = strings.TrimSpace(addr)
		if conn, ok := onlineConns[addr]; ok {
			fmt.Printf("send message [%s] to %s ", sendMessages, addr)
			conn.Write([]byte(sendMessages))
		}
	}
}

func main() {

	listenSocket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	fmt.Println("begin to listen ")
	defer listenSocket.Close()
	go ConsumMessage()
	for {
		conn, err := listenSocket.Accept()
		CheckError(err)
		// 将这个链接对象存储到一个map里面
		addr := fmt.Sprintf("%s", conn.RemoteAddr())
		onlineConns[addr] = conn
		for v, _ := range onlineConns {
			fmt.Println(v)
		}
		go ProcessInfo(conn)
	}
}
