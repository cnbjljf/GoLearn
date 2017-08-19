// main
package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "testconf.ini")
	if err != nil {
		fmt.Println("new config failed,err", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("happend a error:err", err)
	}
	fmt.Println("server-port", port)

	cliPort, err := conf.Int("client::port")
	if err != nil {
		fmt.Println("happend a error:err", err)
	}
	fmt.Println("client port", cliPort)

	log_level := conf.String("logs::log_level")
	if len(log_level) == 0 {
		log_level = "debug"
	}
	fmt.Println("log_level", log_level)

	log_path := conf.String("logs::log_path")
	fmt.Println("log_path", log_path)
}
