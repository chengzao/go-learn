// select来设置超时
package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
