package main

import "fmt"

type Human struct {
	Name string
	age int
}

var str1 string = "abc"

var people = Human{Name: "zhangsan", age: 18}
func main(){
	fmt.Printf("%v\n", people) //{zhangsan}
	fmt.Printf("%+v\n", people) //{Name:zhangsan}
	fmt.Printf("%#v\n", people) //main.Human{Name:"zhangsan"}
	fmt.Printf("%T\n", people) //main.Human
	fmt.Printf("%t\n",true)
	fmt.Printf("%b\n", 5)// 101
	fmt.Printf("%c\n",0x4E2D) //中
	fmt.Printf("%d",0x12) //18
	fmt.Printf("%o\n", 10) //12
	fmt.Printf("%q\n", 0x4E2D) // '中'
	fmt.Printf("%x", 13) // d
	fmt.Printf("%X\n", 13)
	fmt.Printf("%U\n", 0x4E2D) //U+4E2D
	fmt.Printf("%s\n", []byte("Go语言"))
	fmt.Printf("%q\n", "Go语言")
	fmt.Printf("%x\n", "golang")
	fmt.Printf("%X\n", "golang")
	fmt.Printf("%p\n", &people)
	fmt.Printf("%T", str1)
}
