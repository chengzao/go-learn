package main

import "fmt"

// 返回a、b中最大值.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//命名返回值
func ret(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	x := 3
	y := 4
	z := 5

	maxXy := max(x, y) //调用函数max(x, y)
	maxXz := max(x, z) //调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, maxXy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, maxXz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它
	fmt.Println(ret(17))
}
