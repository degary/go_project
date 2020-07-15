package main

import "fmt"

func main()  {
	//1.& 取地址
	n := 18
	p := &n
	fmt.Println(p) //0xc0000180b0
	fmt.Printf("%T\n",p) // *int : int类型的指针

	//2. *: 根据地址取值
	m := *p
	fmt.Println(m) // 18

}
