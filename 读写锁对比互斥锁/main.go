package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//var rwlock sync.RWMutex
var x int
var mutex sync.Mutex

func write() {
	for i := 0; i < 100; i++ {
		//rwlock.Lock()
		mutex.Lock()
		x = x + 1
		time.Sleep(10 * time.Millisecond)
		//rwlock.Unlock()
		mutex.Unlock()
	}
	wg.Done()
}

func read() {
	for i := 0; i < 100; i++ {
		//rwlock.RLock()
		mutex.Lock()
		time.Sleep(time.Millisecond)
		//rwlock.RUnlock()
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixNano()
	wg.Add(1)
	go write()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("cost", cost, "ms")
}
