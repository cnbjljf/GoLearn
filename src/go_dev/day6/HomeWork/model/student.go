// student
package model

import (
	"encoding/json"
	"fmt"
	"go_dev/day6/HomeWork/constValue"
	"io/ioutil"
	"log"
	"os"
)

var stuJson map[int]Student

type Student struct {
	Name      string
	ID        int
	Sex       string
	Age       int
	BrrowBook []Book
}

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
	//	for k, v := range j {
	//		fmt.Printf("ID: %d, student:%v\n", k, v)
	//	}
	fmt.Println("saving student's data successfully!!")
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
		return stuJson, nil
	}
}

func LoginStu(name string, pwd int) (data map[int]Student, ok bool) {
	// Student登录的检查
	data, err := GetStuOldData()
	if err != nil {
		log.Fatalln("happend a error when loading data")
		return data, false
	}
	for _, v := range data {
		if name == v.Name && pwd == v.ID {
			return data, true
		}
	}
	return data, false
}
