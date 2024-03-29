package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	User            // 匿名嵌套
	Password string `json:"password,omitempty"`
}

func OmitPasswordDemo() {
	u1 := User{
		Name:     "七米",
		Password: "123456",
	}

	u2 := PublicUser{
		User:     u1,
		Password: "",
	}
	fmt.Println(u2)
	b, err := json.Marshal(u2)
	// b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b) // str:{"name":"七米"}
	// fmt.Println(u2.User.Password)
}
