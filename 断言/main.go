package main

import "fmt"

type Inter interface {
	Ping()
	Pong()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (s St) Ping() {
	fmt.Println("ping")
}

func (s *St) Pong() {
	fmt.Println("Pong")
}

func main() {
	st := &St{"adnes"}
	var i interface{} = st
	//判断i绑定的实例是否实现了Inter接口
	o := i.(Inter)
	o.Ping()
	o.Pong()

	s := i.(*St)
	fmt.Printf("%s", s.Name)
}
