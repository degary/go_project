package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"11111" b:"33333"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok1 := rt.FieldByName("Name")
	if ok1 {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		fmt.Println("The key of a :", fieldAge.Tag.Get("a"))
		fmt.Println("The key of b :", fieldAge.Tag.Get("b"))
	}
	fmt.Println("type_Name:", rt.Name())
	fmt.Println("type_NumField:", rt.NumField())
	fmt.Println("type_PkgPath:", rt.PkgPath())
	fmt.Println("type_String:", rt.String())
	fmt.Println("type.Kind.String:", rt.Kind().String())
	fmt.Println("type.String()=", rt.String())

	//获取结构类型的字段名称
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name := %v \n", i, rt.Field(i).Name)
	}
	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3, 4, 5)
	sct := reflect.TypeOf(sc)

	//获取slice的元素Type
	scet := sct.Elem()
	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())
	fmt.Println("slice element type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice element type.Name()=", scet.Name())
	fmt.Println("slice element type.PkgPath()=", scet.PkgPath())
	fmt.Println("slice element type.PkgPath()=", sct.PkgPath())
}
