// encodeJson
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"usernameyq"`
	NickName string `json:"nickname"`
	Age      int
	Sex      string
	Email    string
	Phone    string
}

func testStruct() {
	user1 := User{
		UserName: "user1",
		NickName: "Fuck1",
		Age:      19,
		Email:    "1403208@qq.com",
		Phone:    "110001",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("happend a error:", err)
	}
	fmt.Println(string(data))
}

func testInt() {
	var age = 100

	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("happend a error:", err)
	}
	fmt.Println(string(data))
}

func main() {
	//	testStruct()
	testInt()
}
