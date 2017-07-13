// values  存取常量的，用于其他地方调用，意味着只需要修改它就行了,包含一个通用的方法Exist和登录的方法

package constValue

import (
	"os"
)

const (
	BookDataFilePath = "d:/bookData.json"
	StuDataFilePath  = "d:/studentData.json"
	AdminMsg         = `
	1. 添加图书信息
	2. 添加学生信息
	3. 后台管理（删除书籍与学生信息）
	4. 显示当前注册的图书
	5. 显示当前注册的学生信息
	6. 退出`

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

func LoginAdmin(name, pwd string) (role string, ok bool) {
	// admin登录的检查
	if name == "admin" && pwd == "123456" {
		return "admin", true
	}
	return "", false
}
