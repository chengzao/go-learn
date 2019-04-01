package main

import (
	"fmt"
	"time"
)

func main() {
	timeF()
	useF()
}

/*
switch sExpr {
	case expr1:
		some instructions
	case expr2:
		some other instructions
	case expr3:
		some other instructions
	default:
		other code
	}
*/
func baseF() {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
}

//使用fallthrough强制执行后面的case代码
func useF() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func switchF() {
	a := 1
	switch a {
	case 0:
		println("0")
	case 1:
		println("1")
	}
}

func fallthroughF() {
	a := 0
	switch a {
	case 0:
		fallthrough
	case 1:
		test1()
	}
}

func test1() {
	println("test func")
}

func defaultF() {
	i := 2
	switch i {
	case 0:
	case 1:
		println("1")
	default:
		test1() // 当 i 不等于 0 或 1 时调用
	}
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+':
		return true
	}
	return false
}

func timeF() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("today is :", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("saturday is too far away.")
	}
}
