// add
package add

import (
	"fmt"
	_ "go_dev/day2/example1/fortest"
)

var Name string = "Leo"
var Age int = 22

func init() {
	fmt.Println("init in add1")
	fmt.Println("add1.Name", Name)
	fmt.Println("add1.Age", Age)
}
