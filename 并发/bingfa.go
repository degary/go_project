package main

import (
	"fmt"
	"runtime"
	"time"
)

//import (
//	"fmt"
//	"time"
//)
//
//func Hello()  {
//	fmt.Println("Hello Goroutine!")
//}
//
//func main()  {
//	go Hello()
//	fmt.Println("The func Main done")
//	time.Sleep(time.Second)
//}


//var wg sync.WaitGroup
//
//func Hello(x int){
//	defer wg.Done() // goroutine 结束就登记-1
//	fmt.Println("Hello Goroutine!", x)
//}
//
//func main()  {
//	for i :=0;i <10;i++ {
//		wg.Add(1) // 启动一个goroutine就登记+1
//		go Hello(i)
//	}
//	wg.Wait() // 等待所有登记的goroutine都结束
//}

func a(){
	for i := 0; i< 10; i++ {
		//fmt.Println("A: ",i)
	}
}

func b(){
	for i := 0; i< 10; i++ {
		//fmt.Println("B: ",i)
	}
}

func main()  {
	start := time.Now()
	runtime.GOMAXPROCS(1)  // Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。
	go a()
	go b()
	time.Sleep(time.Second)
	end := time.Now()
	fmt.Println(end.Sub(start))
}