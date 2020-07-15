package main

import "fmt"

func main()  {
	var i int = 100
	//输出i的地址
	fmt.Printf("i的地址为%v\n", &i)
	// %p与&i，可正常输出，%p无取址功能
	fmt.Printf("i的地址为%p\n", &i) // i的地址为0xc000014088
	// %p与i，看出输出格式有误，不能输出正确的地址
	fmt.Printf("i的地址为%p\n", i) // i的地址为%!p(int=100)

	// 定义指针变量ptr，类型是*int，值为i的地址
	var ptr *int = &i
	fmt.Printf("ptr的值为%v\n", ptr) // ptr的值为0xc000014088
	// 获取指针类型变量ptr指向的值，用*
	fmt.Printf("ptr存的地址指向的变量的值为%d\n", *ptr) // ptr存的地址指向的变量的值为100
	// 输出ptr自身的地址
	fmt.Printf("ptr的地址为%v\n", &ptr) // ptr的地址为0xc000098020

	// 语法错误，不能用"*ptr2"这种形式的变量名
	// var *ptr2 int = 200
	// fmt.Printf("ptr2的值为%v\n", ptr2) // line24 syntax error: unexpected *, expecting name

	// 示例
	n := 150
	ptr3 := &n
	*ptr3 = 300
	fmt.Printf("n = %v\n", n) // n = 300
}
