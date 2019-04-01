// 并发模式之协程池
package main

import (
	"concurrent/coroutine/namePrinter"
	"concurrent/coroutine/worker"
	"sync"
)

var names = []string{
	"steve",
	"bob",
}

func main() {

	// 1. 启动两个 goroutine，等待执行任务
	p := worker.New(2)

	var wg sync.WaitGroup
	wg.Add(3 * len(names))

	// 2. 创建 worker，扔到 goroutine 池中
	for i := 0; i < 3; i++ {
		for _, namex := range names {
			worker := namePrinter.NamePrinter{
				Name: namex,
			}
			go func() {
				p.Run(&worker)
				wg.Done()
			}()
		}
	}

	// 3. 等待添加任务完毕
	wg.Wait()

	// 4. 关闭 goroutine 池
	p.Shutdown()
}
