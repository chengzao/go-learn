// 接口实现多态
package main

import (
	// 相对于 GOPATH/src 下的地址
	"interface/multi/api"
	"interface/multi/impl"
)

func sendNotification(n api.Notifier) {
	n.Notify()
}

func main() {
	u := impl.User{"nana"}
	sendNotification(&u) // user {nana}

	a := impl.Admin{"zhao"}
	sendNotification(&a) // admin {zhao}
}
