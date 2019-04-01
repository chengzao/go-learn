package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用4个goroutine来完成10个任务
const (
	taskNum      = 10
	goroutineNum = 4
)

var countDownLatch sync.WaitGroup

func worker(name string, taskChannel chan string) {
	defer countDownLatch.Done()
	for {
		// 1. 不断的阻塞等待分配工作
		task, ok := <-taskChannel
		if !ok {
			fmt.Printf("channel-go closed and channel-go is empty\n")
			return
		}

		//fmt.Printf("worker %s start %s\n", name, task)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("worker %s complete %s\n", name, task)
	}
}

func main() {
	countDownLatch.Add(goroutineNum)
	// 1. 创建有缓冲区的string channel-go
	taskChannel := make(chan string, taskNum)

	// 2. 创建 4 个goroutine去干活
	for i := 0; i < goroutineNum; i++ {
		go worker(fmt.Sprintf("worker %d", i), taskChannel)
	}

	// 3. 向通道加入task
	for i := 0; i < taskNum; i++ {
		taskChannel <- fmt.Sprintf("task %d", i)
	}

	// 4. 关闭通道：
	// 当通道关闭后，goroutine 依旧可以从通道接收数据，但是不能再向通道里发送数据。
	// 能够从已经关闭的通道接收数据这一点非常重要，因为这允许通道关闭后依旧能取出其中缓冲的全部值，而不会有数据丢失。
	// 从一个已经关闭且没有数据的通道里获取数据，总会立刻返回，并返回一个通道类型的零值
	close(taskChannel)

	// 5. 等待
	countDownLatch.Wait()
}
