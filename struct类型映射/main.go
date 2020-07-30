package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		UserId   int    `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name"`
	}
	//输出json格式
	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	fmt.Printf("type of j is %T\n", j)

	//获取tag中的值
	t := reflect.TypeOf(u)
	field := t.Elem().Field(1)
	fmt.Println(field.Tag.Get("json"))
}
