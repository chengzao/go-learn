// 无缓冲的 Channel 使用
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

// Channel 完整的类型是 "chan 数据类型"
func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 1. 阻塞等待接球，如果通道关闭，ok返回false
		ball, ok := <-court
		if !ok {
			fmt.Printf("channel-go already closed! Player %s won\n", name)
			return
		}

		random := rand.Intn(100)
		if random%13 == 0 {
			fmt.Printf("Player %s Lose\n", name)
			// 关闭通道
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		// 2. 发球，阻塞等待对方接球
		court <- ball
	}
}

// 两个 player 打网球，即生产者和消费者模式（互为生产者和消费者）
func main() {
	wg.Add(2)

	// 1. 创建一个无缓冲的通道
	// Channel 完整的类型是 "chan 数据类型"
	court := make(chan int)

	// 2. 创建两个 goroutine
	go player("zhangsan", court)
	go player("lisi", court)

	// 3. 发球：向通道发送数据，阻塞等待通道对端接收
	court <- 1

	// 4. 等待输家出现
	wg.Wait()
}
