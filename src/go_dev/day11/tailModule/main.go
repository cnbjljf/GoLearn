// main
package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "./my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,  // C
		Follow:    true,  // 一直读取
		MustExist: false, // 是否要必须存在，如果为false的话那么文件不存在也行，否则为true的话报错
		Poll:      true,  //
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file Close reopen,filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg", msg)
	}
}
