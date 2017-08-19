// 批量set
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c,err := redis.Dial("tcp","192.168.12.23:6379")
	if err != nil {
		fmt.Println("conn redis failed",err)
		return
	}
	
	defer c.Close()
	_ , err = c.Do("MSet","abc",,100)
	
}

