package main

import "fmt"

func main() {
	baseF()
	// useF()
	// useMultiF()
}

func baseF() {
	a := 11
	x := 2
	//变量b的作用域只能在该条件逻辑块内
	if b := 1; a > b {
		fmt.Println("b is ", b)
		fmt.Println("x is ", x)
	} else {
		fmt.Println("b is ", b)
	}
}

func computedValue(x int) int {
	return x
}

//条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，
//其他地方就不起作用

func useF() {
	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := computedValue(2); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	// fmt.Println(x)   //
}

//多个条件
func useMultiF() {
	var a int = 2
	if a == 3 {
		fmt.Println("The a is equal to 3")
	} else if a < 3 {
		fmt.Println("The a is less than 3")
	} else {
		fmt.Println("The a is greater than 3")
	}
}
