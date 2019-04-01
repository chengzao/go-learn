package main

import "fmt"

func main() {
	floatF()
}

// 类型转换
// 表达式 T(v) 将值 v 转换为类型 T
func floatF() {
	var a float32 = 1.4142135
	var b = 1.4142135
	var c = float32(b) //转换float32
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%t \n", a == c)
}
