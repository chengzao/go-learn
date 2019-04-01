package main

import (
	"fmt"
	"math"
)

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

// 枚举
func enums() {
	// 使用 const 块来实现枚举
	const (
		java = 0
		cpp  = 1
		c    = 2
	)
	fmt.Println(java, cpp, c) // 0 1 2
	// 使用 iota 块来实现自增枚举
	const (
		java1 = iota
		cpp1
		c1
	)
	fmt.Println(java1, cpp1, c1) // 0 1 2
}

type ByteSize float64

func iota_byte() {
	const (
		_           = iota             // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
		EB                             // 1 << (10*6)
		ZB                             // 1 << (10*7)
		YB                             // 1 << (10*8)
	)
	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)
	fmt.Println("TB", TB)
	fmt.Println("PB", PB)
	fmt.Println("EB", EB)
	fmt.Println("ZB", ZB)
	fmt.Println("YB", YB)
}

func week() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

func main() {
	consts()
	enums()
	iota_byte()
	week()
}
