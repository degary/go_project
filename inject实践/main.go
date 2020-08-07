package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name=%s,company=%s,level=%s,age=%d!\n", name, company, level, age)
}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func main() {
	/*
	   以下 实现了对函数的参数注入

	   	//控制实例创建
	   	inj := inject.New()

	   	//实参注入
	   	inj.Map("tom")
	   	inj.MapTo("tencent",(*S1)(nil))
	   	inj.MapTo("T4",(*S2)(nil))
	   	inj.Map(23)

	   	//函数反转调用
	   	inj.Invoke(Format)
	*/
	s := Staff{}

	//控制实例的创建
	inj := inject.New()

	//初始化注入值
	inj.Map("Tome")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(23)

	//实现对struct的注入
	inj.Apply(&s)

	//打印
	fmt.Printf("s=%v\n", s)

}
