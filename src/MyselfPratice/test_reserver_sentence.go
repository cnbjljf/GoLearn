package main

import (
	"fmt"
	"go_dev/reversal"
)

func main() {
	r := reversal.ReserverSentence("!上 1海 自 来 水 来 自 海 上@")
	fmt.Println(r)
}
