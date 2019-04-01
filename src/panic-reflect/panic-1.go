// panic & recover
package main

import (
	"errors"
	"fmt"
)

func recove() {
	defer func() {
		// func recover() interface{}，表示 recover() 函数的返回类型可以是各种类型，所以要判断是否是 error
		// 使用 recover() catch panic，防止程序直接退出
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println(err) // runtime error: integer divide by zero
		} else {
			panic(errors.New("not known error"))
		}
	}()

	b := 0
	a := 5 / b // panic: runtime error: integer divide by zero
	fmt.Println(a)

	//panic("123") // panic: not known error
}

func main() {
	recove()
}
