// parseConf
package main

import (
	"flag"
	"fmt"
)

func main() {
	var configPath string
	var logLevel int
	flag.StringVar(&configPath, "c", "", "input a config file path")
	flag.IntVar(&logLevel, "d", 10, "input log level")

	flag.Parse()
	fmt.Println("path:", configPath)
	fmt.Println("log level", logLevel)
}
