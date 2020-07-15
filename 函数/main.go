package main

import "fmt"

//函数的定义
//命名的返回值就相当于在函数中已经声明了变量
func sum(x int, y int)(ret int){
	ret =  x + y
	return  //使用命名返回值,return后可以省略
}

//没有返回值的函数
func f1(x int, y int)  {
	fmt.Println(x + y)

}

//没有参数,但是有返回值
func f2() int {
	return 3
}

//有多个返回值
func f3()(int, int)  {
	return 1, 2
}

//当参数中连续两个参数的类型一致时,可以简写
func f4(x,y int)  int{
	return x + y
}

//可变长参数
func f5(x string, y ...int){
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("Type y: %T\n",y) //[]int
}

func main() {
	r := sum(10, 20)
	fmt.Println(r)
	f5("a",1,2,4,5)
}