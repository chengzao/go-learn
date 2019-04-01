// 闭包
package main

import "fmt"

// 该函数的返回值是一个函数：func(value int) int
func adder() func(value int) int {
	sum := 0 // 自由变量，Go 的闭包是可以含有自由变量的
	return func(value int) int {
		sum += value
		return sum
	}
}

func main() {
	add := adder() // 此处返回一个函数
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d", i, add(i)) //调用函数add
		fmt.Println()
	}
}
