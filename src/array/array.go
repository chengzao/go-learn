package main

import "fmt"

func defineArray() [3]int {
	// 定义数组，不赋初值（使用默认值）
	var arr1 [5]int // [0 0 0 0 0]
	// 定义数组，赋初值
	arr2 := [3]int{1, 2, 3} // [1 2 3]
	// 定义数组，由编译器来计算长度，不可写成[]，不带长度或者 ... 的表示切片
	arr3 := [...]int{4, 5, 6, 7} // [4 5 6 7]
	// 创建指针数组
	arr4 := [2]*string{new(string), new(string)}
	*arr4[0] = "hello"
	*arr4[1] = "go"
	// 为指定索引位置设置值
	arr5 := [3]int{1: 10} // [0,10,0]
	// 二维数组
	var grid [4][5]int // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
	// 数组拷贝，直接复制一份 arr2 给 arr6
	arr6 := arr2

	fmt.Println("arr1", arr1)
	fmt.Println("arr2", arr2)         //
	fmt.Println("arr3", arr3)         //
	fmt.Println("arr4", arr4)         // [0xc00000e1e0 0xc00000e1f0]
	fmt.Println("arr5", arr5)         //
	fmt.Println("arr6", arr6)         //
	fmt.Println("grid", grid)         //
	fmt.Println("*arr4[0]", *arr4[0]) // hello
	return arr2
}

// 数组是值传递，这里的入参会拷贝一份数组（使用指针可以实现"引用传递"）
func loopArray(arr2 [3]int) {
	// 通用方法
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	// 最简方法，只获取数组下标
	for i := range arr2 {
		fmt.Println(arr2[i])
	}
	// 最简方法，获取数组下标和对应的值
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	// 最简方法，只获取值，使用 _ 省略变量
	for _, v := range arr2 {
		fmt.Println(v)
	}
}

func arr1() {
	var arr [10]int //
	arr[0] = 42
	fmt.Printf("The first arr1 is %d\n", arr[0])
}

func rangeF() {
	list := []string{"a", "b", "c"}
	for k, v := range list {
		println(k, v)
	}
}

func charF() {
	for pos, char := range "aux" {
		fmt.Printf("character '%c' starts at byte position%d \n", char, pos)
	}
}

func main() {
	defineArray()

	//arr := [3]int{1, 2, 3}
	//loopArray(arr)
}
