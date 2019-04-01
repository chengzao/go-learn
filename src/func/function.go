package main

import (
	"fmt"
	"math"
)

var a string

func main() {
	// rec(1)
	a = "abc"
	println("main " + a) // main abc
	f()
	funcTest1()
	myfunc(1, 2)
	test1()
	test2()
}

// 递归
func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d \n", i)
}

func f() {
	a := "def"
	println("f " + a) //f def
	g()
}
func g() {
	println("g " + a) // a abc
}

func funcTest1() {
	a := func() { //← 定义一个匿名函数，并且赋值给 a
		println("Hello")
	} //← 这里没有 ()
	a() //← 调用函数
}

// 变参:不定参数
func myfunc(args ...int) {
	for _, n := range args {
		fmt.Printf("And the number is: %d\n", n)
	}
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) sum() float64 {
	return v.X + v.Y
}

func test1() {
	v := &Vertex{3, 4}
	fmt.Println(v.sum())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func test2() {
	f := MyFloat(-math.Sqrt(2))
	fmt.Println(f.Abs())
}
