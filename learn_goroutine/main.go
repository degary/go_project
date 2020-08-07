package main

import (
	"fmt"
	"runtime"
	"time"
)

//func Hello(i int)  {
//	fmt.Println("Hello world",i)
//}
//
//func main()  {
//	for i :=0; i<10;i++ {
//		go Hello(i)
//	}
//	time.Sleep(100*time.Millisecond)
//	fmt.Println("main thread terminate")
//}

func number() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	go number()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
}
