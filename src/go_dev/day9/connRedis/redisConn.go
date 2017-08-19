// redisConn
package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.12.23:6379")
	if err != nil {
		fmt.Println("conn redis happend a err", err)
	}
	defer c.Close()
	_, err = c.Do("set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed", err)
		return
	}
	fmt.Println(r)
}
