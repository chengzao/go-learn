package main

import (
	"fmt"
	"hello/utils"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Printf("golang mod ! \n")
	utils.PrintText("from hello mod utils \n")
	beego.Run()
}
