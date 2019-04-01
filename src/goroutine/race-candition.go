// 锁机制: 原子类 atomic
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
)

func incCount(int) {
	defer waitGroup.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// 当前的 goroutine 主动让出资源，从线程退出，并放回到队列，
		// 让其他 goroutine 进行执行
		runtime.Gosched()
		value++
		counter = value
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	waitGroup.Add(2)

	go incCount(1)
	go incCount(2)

	waitGroup.Wait()
	fmt.Println(counter) // 正确为4，实际上为2
}
