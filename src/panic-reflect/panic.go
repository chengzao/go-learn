// panic 和 recover
package main

import "fmt"

func main() {
	fmt.Printf("hello world my name is %s, I'm %d\r\n", "xxxx", 26)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	myPainc() // 出错函数
	// myPanic2() // 正确函数
	fmt.Printf("我是正确执行到这里")
}
func myPainc() {
	var x = 30
	var y = 0
	var c = x / y
	// panic("我一个大错误！")
	fmt.Println(c)
}

func myPanic2() {
	var x = 30
	var y = 1
	var c = x / y
	// panic("我一个大错误！")
	fmt.Println(c)
}
