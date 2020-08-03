package main

import "fmt"

//chain 函数的输入参数和输出参数类型相同,都是 chan int
//chain函数的功能是将chain 内的数据统一加1

func chain(in chan int) chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	out := chain(chain(chain(in)))
	for v := range out {
		fmt.Println(v)
	}
}
