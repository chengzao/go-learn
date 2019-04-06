package main

import (
	"fmt"
	"math"
	"math/cmplx"
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

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func calcTriangle(a, b int) int {
	var c int
	// 类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//声明未赋值
var str0 string
var str1 string = "str1"
var str2 = "str2"

//str3 := "str3"   //错误
var str3 = "tstr3-1"

func varF() {
	//var str5 int //声明str5后需要使用,否侧报错
	str4 := "str4" //声明str4赋值str4后需要使用,否侧报错
	println(str0)  //
	println(str4)  // "str4"
}

func test() {
	str3 = "str3"
	println(str1) // "str1"
	println(str2) // "str2"
	println(str3) // "str3"
}

func utfF() {
	println('a') // 97
	println("a") // "a"
}

func runeByteF() {
	s := "Go编程"
	println("'s' 的len ==>", len(s))                //8
	println("'编'的len ==>", len(string(rune('编')))) //3
	println("rune的len ==>", len([]rune(s)))        //4
}

func main() {
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	triangle()

	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
	fmt.Println(aa, bb)

	euler()
}
