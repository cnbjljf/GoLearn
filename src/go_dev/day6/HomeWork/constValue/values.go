// values  存取常量的，用于其他地方调用，意味着只需要修改它就行了,包含一个通用的方法Exist和登录的方法

package constValue

import (
	"fmt"
	"os"
	"bufio"
)

const (
	BookDataFilePath = "d:/bookData.json"
	StuDataFilePath  = "d:/studentData.json"
	AdminMsg := `
	1. 添加图书信息
	2. 添加学生信息
	3. 借书
	4. 后台管理（删除书籍与学生信息）
	5. 显示当前注册的图书
	6. 显示当前注册的学生信息
	7. 显示当前注册的学生信息`
	
	
	CustonmerMsg = `
	1. 借书
	2. 还书
	3. 退出
	`
)

func Exist(filename string) bool { // 判断指定文件是否存在的
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}


func LoginAdmin() (role string,bool){
	// admin登录的检查
	var name string
	var pwd string
	fmt.Println("请输入用户名")
	fmt.Scanln(&name)
	fmt.Println("请输入密码")
	fmt.Scanln(&pwd)
	
	if name == "admin" && pwd == "123456"{
		return "admin",true
	}
	return "",false
}
