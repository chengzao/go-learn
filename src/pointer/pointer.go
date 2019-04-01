package pointer

import "fmt"

// 值传递，函数参数拷贝了一份外界的 a, b
func swap_by_value(a, b int) {
	b, a = a, b
}

// Go 只有值传递，想实现引用传递，使用指针
// *int 代表是指针类型，此时会将外界传入的 &a 拷贝给 这里的a，即这里的 a 拿到的是外界的 a 的地址
// 通过 *a，由于 a 是 &a，这里的 *a 相当于 *(&a) ，即从地址中取值
// 由于函数内部直接操作的是外界的 a,b 的内存地址，所以可以实现引用传递
func swap_by_pointer(a, b *int) {
	*b, *a = *a, *b
}

func main() {
	// use
	a, b := 3, 4
	swap_by_value(a, b)
	fmt.Println(*(&a), b) // 3 4 没有实现交换
	fmt.Println("\n")

	// use
	swap_by_pointer(&a, &b)
	fmt.Println(a, b)
}
