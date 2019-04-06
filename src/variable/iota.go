package main

import "fmt"

// iota枚举
func constF() {
	// iota
	const (
		a           = iota       //0
		b           = "abc"      //"abc"
		c, d        = iota, iota //2 2
		e           = iota       //3
		f    string = "fff"      //"fff"
	)

	println("a ==>", a)
	println("b ==>", b)
	println("c ==>", c)
	println("d ==>", d)
	println("e ==>", e)
	println("f ==>", f)
}

func enumsF() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
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

// iota
func init() {
	const (
		a = iota
		b
		c
	)
	println("init a b c:", a, b, c)
}

func main() {
	constF()
	enumsF()
}
