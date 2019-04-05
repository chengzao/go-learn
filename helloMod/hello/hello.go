package main

import (
	"fmt"
	"hello/utils"
	"importMod/hello"

	"github.com/jinzhu/configor"
)

// run : go run hello.go
func main() {
	fmt.Println("Hello Mod")
	fmt.Println("使用外部包", configor.Config{})

	fmt.Println("使用项目内包hello", utils.Hello("Hello"))

	hello.Hello("引用另外一个本地的外部包importMod")
}
