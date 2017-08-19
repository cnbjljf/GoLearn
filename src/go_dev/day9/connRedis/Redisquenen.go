// Redisquenen
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
	_,err = c.Do("lpush","book list","abc","ceg",300)
	if err != nil {
		fmt.Println(err)
		return
	}
	r,err := redis.String(c.Do("loop","book list"))
	if err != nil {
		fmt.Println("happend a err",err)
	}
	fmt.Println(r)
