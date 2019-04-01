package main

import (
	"fmt"
)

// user 类
type user struct {
	name string
	mail string
}

// 定义接口

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	// 接口方法
	notify()
}

// notifier2 是一个定义了通知类行为的接口
type notifier2 interface {
	// 接口方法
	notify2()
}

// 实现接口: 使用 值 接收者实现接口

// notify 是使用值接收者实现 notifier interface 接口的方法
// sendNotification(&u) 和 sendNotification(u) 都可
func (u user) notify() {
	fmt.Println("notify", u)
}

// notify2 是使用指针接收者实现 notifier2 interface 接口的方法
// 只能使用 sendNotification2(&u)
func (u *user) notify2() {
	fmt.Println("notify2", *u)
}

// sendNotification 接受一个实现了 notifier 接口的值并发送通知
func sendNotification(n notifier) {
	n.notify()
}

// sendNotification2 接受一个实现了 notifier2 接口的值并发送通知
func sendNotification2(n notifier2) {
	n.notify2()
}

//值接收者（func (u User) Notify() error）操作的是 User 的副本（不管调用者使用值还是地址），
//所以不会改变其内部的值，这里输出还是 Damon

//指针接收者（func (u *User) Notify() error）操作的是 User 本身，
// 所以其内部的值会发生变化，这里输出是 alimon

// 使用接口
func main() {
	u := user{"golang", "golang@golang.org"}
	sendNotification(u)  // notify {golang golang@golang.org}
	sendNotification(&u) // notify {golang golang@golang.org}

	u2 := user{"javascript", "javascript@javascript.org"}
	sendNotification2(&u2) // notify2 {javascript javascript@javascript.org}
}
