// values  存取常量的，用于其他地方调用，意味着只需要修改它就行了,包含一个通用的方法Exist和登录的方法

package constValue

import (
	"fmt"
	"os"
	"time"
)

const (
	BookDataFilePath = "d:/bookData.json"
	StuDataFilePath  = "d:/studentData.json"
	LogFile          = "d:/BookManager.log"
	AdminMsg         = `
	1. 添加图书信息
	2. 添加学生信息
	3. 后台管理（删除书籍与学生信息）
	4. 显示当前注册的学生信息
	5. 显示当前注册的图书
	6. 退出`

	CustonmerMsg = `
	1. 借书
	2. 还书
	3. 退出
	`
	ManageChoice = `
	1. 管理学生信息
	2. 管理图书信息
	`
	ActionChoice = `
	1. 删除信息
	2. 修改信息
	`
)

func Exist(filename string) bool { // 判断指定文件是否存在的
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func Logger(username, action string, content interface{}) {
	ct, _ := content.(string)
	f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("opening file happend a error!,", err)
		return
	}
	logContent := fmt.Sprintf("%s user:[%s] action:[%s] info:[%s]\n", time.Now().Format("2006-01-02 15:04:05"),
		username, action, ct)
	f.WriteString(logContent)
	defer f.Close()
}

func LoginAdmin(name, pwd string) (role string, ok bool) {
	// admin登录的检查
	if name == "admin" && pwd == "111" {

		return "admin", true
	}
	return "", false
}
