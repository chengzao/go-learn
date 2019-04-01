package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown  int64
	waitGroup sync.WaitGroup
)

func doWork(name string) {
	defer waitGroup.Done()
	for {
		time.Sleep(250 * time.Millisecond)
		// 记载关机标志
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Println("shutDown, ", name)
			break
		}
	}
}

func main() {
	waitGroup.Add(2)

	go doWork("A")
	go doWork("B")

	// 给定goroutine执行的时间
	time.Sleep(1000 * time.Millisecond)

	// 设定关机标志
	atomic.StoreInt64(&shutdown, 1)

	waitGroup.Wait()
}
