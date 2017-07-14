//
package main

import (
	//	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"os"
)

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

type stu struct {
	Name  string
	bbook []book
}

type book struct {
	Name string
}

func main() {

}
