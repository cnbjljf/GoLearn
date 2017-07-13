// values  存取常量的，用于其他地方调用，意味着只需要修改它就行了

package constValue

import (
	"os"
)

const (
	BookDataFilePath = "d:/bookData.json"
	StuDataFilePath  = "d:/studentData.json"
)

func Exist(filename string) bool { // 判断指定文件是否存在的
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
