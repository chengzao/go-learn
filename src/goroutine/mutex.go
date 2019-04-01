package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// 两个 goroutine 同时操作的变量，竞态变量
	counter   int
	waitGroup sync.WaitGroup
	// 锁，定义一段临界区
	lock sync.Mutex
)

func incCount(int) {
	defer waitGroup.Done()
	for count := 0; count < 2; count++ {
		lock.Lock()
		{ // Lock() 与 UnLock() 之间的代码都属于临界区，{}是可以省略的，加上看起来清晰
			value := counter
			// 当前的 goroutine 主动让出资源，从线程退出，并放回到队列，
			// 让其他 goroutine 进行执行
			// 但是因为锁没有释放，调度器还会继续安排执行该 goroutine
			runtime.Gosched()
			value++
			counter = value
		}
		lock.Unlock()
		// 释放锁，允许其他正在等待的 goroutine 进入临界区
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	waitGroup.Add(2)

	go incCount(1)
	go incCount(2)

	waitGroup.Wait()
	fmt.Println(counter) // 4
}
