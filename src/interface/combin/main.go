// 组合接口
package main

import (
	// 相对于 GOPATH/src 下的地址
	"interface/combin/api"
	"interface/combin/impl"
)

func sendNotification(n api.Api12) {
	n.Api1() // api1 nana
	n.Api2() // api2 nana
}

func main() {
	api12 := impl.Api12Impl{"nana"}
	sendNotification(&api12)
}
