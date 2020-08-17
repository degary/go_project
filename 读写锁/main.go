package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwlock sync.RWMutex
var x int

func write() {
	rwlock.Lock()
	fmt.Println("write lock")
	x = x + 1
	time.Sleep(10 * time.Second)
	fmt.Println("wtite unlock")
	rwlock.Unlock()
	wg.Done()
}

func read(i int) {
	fmt.Println("wait for write lock")
	rwlock.RLock()
	fmt.Printf("goroutine:%d,x=%d\n", i, x)
	//time.Sleep(time.Second*1)
	rwlock.RUnlock()
	wg.Done()
}

func main() {
	wg.Add(1)
	go write()
	time.Sleep(5 * time.Millisecond)
	wg.Done()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()

}
