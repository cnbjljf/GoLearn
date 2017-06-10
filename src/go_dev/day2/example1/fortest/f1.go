// f1
package fortest

import (
	"fmt"
)

var Age int = 111
var Name string = "Leo"

func init() {
	fmt.Println("init in f1")
	fmt.Println("f1.Age", Age)
	fmt.Println("f1.Name", Name)
	Age = 22
	fmt.Println("f1.Age after change", Age)
}
