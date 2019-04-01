package _func

import "fmt"

//======================== 可以有多个返回值 =========================
// 函数可返回多个值
// 接收：q, r := div(10, 3)
// 如果只用其中一个值，另一个用下划线：q, _ := div(10, 3)
func div(a, b int) (int, int) {
	return a / b, a % b
}

//======================== 函数的参数类型可以是 func =========================
// 可以使用函数作为参数，函数参数与内部参数一样，函数名在前，函数类型在后
// 后续传参，可以使用匿名内部函数，也可以先定义函数再传入
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

//======================== 支持可变长参数 =========================
// 可变长参数
func sum2(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}

func main() {
	// use 可以有多个返回值
	q, _ := div(10, 3)
	fmt.Println(q, "\n")

	// use 函数的参数类型可以是 func
	result := apply(func(x int, y int) int {
		return x + y
	}, 10, 4)
	fmt.Println(result, "\n")

	// use 支持可变长参数
	sum := sum2(1, 2, 3)
	fmt.Println(sum, "\n")
}
