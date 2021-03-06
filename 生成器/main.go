package main

import (
	"fmt"
	"math/rand"
)

//done 接收通知推出信号
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)

	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

//done 接收通知推出信号
func GenerataIntB(done chan struct{}) chan int {
	ch := make(chan int, 10)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

//通过select执行扇入(Fan In)操作
func GenerateInt(done chan struct{}) chan int {
	ch := make(chan int)
	send := make(chan struct{})
	go func() {
	Lable:
		for {
			select {
			case ch <- <-GenerateIntA(send):
			case ch <- <-GenerataIntB(send):
			case <-done:
				break Lable
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	//创建一个接收推出信号的chan
	done := make(chan struct{})

	//启动生成器
	ch := GenerateInt(done)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	//通知生产者停止生产
	done <- struct{}{}
	fmt.Println("stop gernarate")

}
