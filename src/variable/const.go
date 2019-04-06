package main

import (
	"fmt"
	"math"
)

func constsF() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

// 常量
func consts() {
	// 指定类型
	const filename string = "filename-const"
	// 不指定类型，表示类型不定
	const a, b = 3, 4
	var c int
	// 由于类型不定，所以这里不需要强转，如果定义为 const a, b int = 3, 4，则需要强转
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, a, b, c)
}

func main() {
	constsF()
	consts()
}
