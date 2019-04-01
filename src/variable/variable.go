package main

import (
	"fmt"
	"math"
)

// 定义包内变量
var (
	aa = 1
	bb = "tianya"
)

// 定义变量，只使用默认初值
func variableZeroValue() {
	var a int    // 0
	var s string // ""
	var b bool   // false
	var c [1]string
	fmt.Println(a, s, b, c)
}

// 定义变量，赋初值
func variableInitialValue() {
	var a, b int = 1, 2
	b = 3
	var s string = "haha"
	fmt.Println(a, b, s)
}

// 类型推断
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

// 最简定义变量方式
func variableShorter() {
	a, b, c, d := 3, 2, true, "hi"
	fmt.Println(a, b, c, d)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	// float64 和 int 可以不加小括号，也可以加上
	// 开方内建函数定义：func Sqrt(x float64) float64
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	triangle()
	fmt.Println(aa, bb)
}
