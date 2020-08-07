package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) String() {
	fmt.Println("User:", this.Id, this.Name, this.Age)
}

func Info(o interface{}) {
	//获取value信息
	v := reflect.ValueOf(o)
	t := v.Type()
	//类型名称
	fmt.Println("Type:", t.Name())

	//访问接口字段名,字段类型和字段信息
	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		//查询类型
		switch value := value.(type) {
		case int:
			fmt.Printf(" %6s: %v = %d\n", field.Name, field.Type, value)
		case string:
			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)
		default:
			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)

		}
	}
}

func main() {
	u := User{1, "Tom", 30}
	Info(u)
}
