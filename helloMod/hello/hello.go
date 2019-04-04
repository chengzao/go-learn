package main

import (
	"fmt"
	"hello/utils"

	"github.com/jinzhu/configor"
)

func main() {
	fmt.Println("Hello Mod")
	fmt.Println("使用外部包", configor.Config{})

	fmt.Println("使用项目内包", utils.Hello("Hello"))
}
