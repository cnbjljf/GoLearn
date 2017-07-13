// student
package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_dev/day6/HomeWork/constValue"
	"io/ioutil"
	"log"
	"os"
)

var (
	LoadOldDataError = errors.New("load the student's data happend a error")
)

type Student struct {
	Name string
	ID   int
	Sex  string
	Age  int
}

var stuJson map[int]Student

func SaveStuData(j map[int]Student) bool {
	// 保存学生信息到文本内容里面
	f, err := os.Create(constValue.StuDataFilePath)
	if err != nil {
		log.Fatalln("Create a student's file happed a error!", err)
		return false
	}
	stuJson, err := json.Marshal(j)
	if err != nil {
		log.Fatalln("format  student's info to json happed a error!", err)
		return false
	}
	f.Write(stuJson)
	f.Sync()
	defer f.Close()
	for k, v := range j {
		fmt.Printf("ID: %d, student:%v\n", k, v)
	}
	fmt.Println("save student's data successfully!!")
	return true
}

func DelStu(j map[int]Student, name string) bool {
	return false
}

func GetStuOldData() (stuJson map[int]Student, err error) {
	// 获取老数据
	stuJson = make(map[int]Student)
	if constValue.Exist(constValue.StuDataFilePath) { // 先判断文件是否存在，存在才读取老文件
		oldF, err := ioutil.ReadFile(constValue.StuDataFilePath) // 读取文本数据
		if err != nil {
			log.Fatalln("read file happend a error:", err)
			return stuJson, LoadOldDataError
		}
		fmt.Println("find the  student's  data file and then ready to load it!!these are exist student!")
		json.Unmarshal(oldF, &stuJson) // 加载之前的数据
//		for k, v := range stuJson {
//			fmt.Printf("ID: %d, student:%v\n", k, v)
//		}
		return stuJson, nil
	} else {
		return stuJson, LoadOldDataError
	}
}


func LoginStu() (role string,bool){
	// Student登录的检查
	var name string
	var pwd string
	fmt.Println("请输入用户名")
	fmt.Scanln(&name)
	fmt.Println("请输入ID")
	fmt.Scanln(&pwd)
	data,err := GetStuOldData()
	if err !=nil {
		log.Fatalln("happend a error when loading data")
	}
	for _,v := range {
		if name == v["Name"] && pwd == v["ID"] {
			return "student",true
		}
	}
	return "",false
}
