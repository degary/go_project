package main

import "fmt"
//Go中 函数的return不是原子操作,在底层是分为两步来执行
//1.返回值赋值
//2.真正的return返回
//函数中如果存在defer函数,那么defer执行的 时机是在1和2之间
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}