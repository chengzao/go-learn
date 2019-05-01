package main

import (
	"fmt"
	"hello/utils"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Printf("golang mod ! \n")
	utils.PrintText("I from helloMod/utils file tool.go \n")
	beego.Run()
}
