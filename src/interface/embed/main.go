// 接口: 嵌入类型内部接口实现提升
package main

import (
	// 相对于 GOPATH/src 下的地址
	"interface/embed/api"
	"interface/embed/impl"
)

func sendNotification(n api.Notifier) {
	n.Notify()
}

func main() {
	u := impl.User{"nana"}
	sendNotification(&u) // user {nana}

	a := impl.Admin{u, "zhao"}
	sendNotification(&a) // user {nana}
}
