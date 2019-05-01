package main

import (
	"fmt"
	"hello/utils"
	"importMod/hello"

	"github.com/jinzhu/configor"
)

// run : go run hello.go
func main() {
	fmt.Println("HelloMod/hello : hello.go")
	fmt.Println("使用外部包 github包: ", configor.Config{})

	fmt.Println("使用项目内包helloMod/utils tools.go: ", utils.Hello("Hello"))

	hello.Hello("exportMod/other.go : 引用另外一个本地的外部包importMod")
}
