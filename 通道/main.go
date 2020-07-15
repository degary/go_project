package main

import (
	"fmt"
	"sync"
	"time"
)

//import "fmt"
//
//// channel 练习
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	// 开启goroutine将0~100的数发送到ch1中
//	go func() {
//		for i := 0; i < 100; i++ {
//			ch1 <- i
//		}
//		close(ch1)
//	}()
//	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//	go func() {
//		for {
//			i, ok := <-ch1 // 通道关闭后再取值ok=false
//			if !ok {
//				break
//			}
//			ch2 <- i * i
//		}
//		close(ch2)
//	}()
//	// 在主goroutine中从ch2中接收值打印
//	for i := range ch2 { // 通道关闭后会退出for range循环
//		fmt.Println(i)
//	}
//}

//func worker(id int, jobs <-chan int, results chan<- int) {
//	for j := range jobs {
//		fmt.Printf("worker:%d start job:%d\n", id, j)
//		time.Sleep(time.Second)
//		fmt.Printf("worker:%d end job:%d\n", id, j)
//		results <- j * 2
//	}
//}
//
//
//func main() {
//	jobs := make(chan int, 100)
//	results := make(chan int, 100)
//	// 开启3个goroutine
//	for w := 1; w <= 3; w++ {
//		go worker(w, jobs, results)
//	}
//	// 5个任务
//	for j := 1; j <= 5; j++ {
//		jobs <- j
//	}
//	close(jobs)
//	// 输出结果
//	for a := 1; a <= 5; a++ {
//		<-results
//	}
//}

//var x int64
//var wg sync.WaitGroup
//
//func add() {
//	for i := 0; i < 5000; i++ {
//
//		x = x + 1
//
//	}
//	wg.Done()
//}
//func main() {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}