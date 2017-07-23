// student
package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_dev/day6/HomeWork/constValue"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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
		//fmt.Println("find the  student's  data file and then ready to load it!!these are exist student!")
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

func deleteStu(username, stuname string) {
	// 删除指定学生名字的信息

	stuInfo, _ := GetStuOldData()
	stuJson = make(map[int]Student)
	var i int
	for _, item := range stuInfo {
		if item.Name == stuname {
			constValue.Logger(username, "delete a student", "the student ["+stuname+"]was deleted!")
			continue
		} else {
			stuJson[i] = item
			i++
		}
	}
	SaveStuData(stuJson)
}

func ManageStudent(username string) {
	// 管理学生信息的功能，包含删除与修改
	for {
		stuInfo, _ := GetStuOldData()
		for _, item := range stuInfo {
			fmt.Printf("name: %-8s ,id: %-5d,sex: %-10s,age: %-2d\n", item.Name, item.ID, item.Sex, item.Age)
		}
		fmt.Println("请输入学生名字,输入quit退出！")
		inputer := bufio.NewReader(os.Stdin)
		result, _, _ := inputer.ReadLine()
		stukName := strings.TrimSpace(string(result))

		if strings.ToLower(stukName) == "quit" {
			return
		}

		var flag bool
		for _, item := range stuInfo {
			if stukName == item.Name {
				flag = true
			}
		}
		if !flag {
			fmt.Println("无效的学生名，没有找到这个人")
			continue
		}

		fmt.Println()
		fmt.Println(constValue.ActionChoice)
		inputer = bufio.NewReader(os.Stdin)
		result, _, _ = inputer.ReadLine()

		action, _ := strconv.Atoi(strings.TrimSpace(string(result)))

		switch action {
		case 1:
			deleteStu(username, stukName)
		case 2:
			fmt.Println("请输入新的信息ID,性别，年龄，必须以逗号分隔！")
			inputer := bufio.NewReader(os.Stdin)
			result, _, _ := inputer.ReadLine()
			allInfo := strings.Split(strings.TrimSpace(string(result)), ",")
			id, _ := strconv.Atoi(allInfo[0])
			sex := allInfo[1]
			age, _ := strconv.Atoi(allInfo[2])

			for k, v := range stuInfo {
				if v.Name == stukName {
					v.ID = id
					v.Sex = sex
					v.Age = age
					stuInfo[k] = v
					constValue.Logger(username, "modifing student's info", stuInfo[k])
					SaveStuData(stuInfo)
				}
			} // end for
		}
	}
}
