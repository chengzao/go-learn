package main

import "fmt"

// create
func createSlice() {
	// 1、使用make函数创建一个字符串切片，长度和容量都是5
	slice1 := make([]string, 5)
	fmt.Println("slice1", slice1)
	// 2、创建一个int切片，长度是3，容量是5
	slice2 := make([]int, 3, 5)
	fmt.Println("slice2", slice2, "len", len(slice2), "cap", cap(slice2))
	// 3、使用字面量创建切片，长度是3，容量是3
	slice3 := []int{1, 2, 3} // [3]int{1, 2, 3}
	fmt.Println("slice3", slice3)
	// 4、创建 nil 切片，长度为0，容量为0
	var slice4 []int
	fmt.Println("slice4", slice4, "len", len(slice4), "cap", cap(slice4))
	// 5、创建空切片，长度为0，容量为0
	slice5 := make([]int, 0)
	fmt.Println("slice5", slice5, "len", len(slice5), "cap", cap(slice5))
	slice6 := []int{}
	fmt.Println("slice6", slice6, "len", len(slice6), "cap", cap(slice6))
	// 6、自定义底层数组，通过该底层数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	// 数组转化为切片，左闭右开 [arr[2]~arr[4])
	slice7 := arr[2:4] // [3,4]
	fmt.Println("slice7", slice7)
	slice8 := arr[2:] // [3,4,5]
	fmt.Println("slice8", slice8)
	slice9 := arr[:4] // [1,2,3,4]
	fmt.Println("slice9", slice9)
	slice10 := arr[:] // [1,2,3,4,5]
	fmt.Println("slice10", slice10)
}

// Use
func useSlice() {
	slice1 := []int{1, 2, 3, 4, 5}
	// 1、根据索引获取切片元素
	fmt.Println("slice[1]", slice1[1]) // 2
	// 2、根据索引修改切片元素
	slice1[3] = 400
	fmt.Println("slice1", slice1) // [1, 2, 3, 400, 5]
	// 3、根据切片创建切片，和根据自定义数组创建切片方式相同，长度是2=3-1，容量是4=5-1
	// 但是需要格外注意，新生成的切片 slice2 和原始切片 slice1 的指针元素指向了相同的底层数组，所以修改元素要注意
	slice2 := slice1[1:3] // [2, 3]
	slice2[1] = 300
	fmt.Println("slice2", slice2) // [2, 300]
	fmt.Println("slice1", slice1) // [1, 2, 300, 400, 5] slice1也发生了变化
	// 4、拷贝 slice 中的元素
	fmt.Println("copy")
	slice3 := []int{0, 0, 0, 0, 0}
	slice4 := []int{1, 2, 3}
	copy(slice3, slice4)
	fmt.Println("slice3", slice3) // [1, 2, 3, 0, 0]
	fmt.Println("slice4", slice4) // [1, 2, 3]
	// 5、删除 slice 中的元素，删除slice5[2]=3
	fmt.Println("delete")
	slice5 := []int{1, 2, 3, 4}
	slice5 = append(slice5[:2], slice5[3:]...)
	fmt.Println("slice5", slice5) // [1, 2, 4]
}

// append 增加切片长度
func appendSlice() {
	// 1、创建原始切片，长度是5，容量是5
	slice := []int{10, 20, 30, 40, 50}
	// 2、reslice 新切片，长度是2，容量是4
	newSlice := slice[1:3] // [20, 30]
	// 由于底层数组还有容量，可以直接追加元素而容量不变
	newSlice = append(newSlice, 60)   // [20, 30 ,60] 长度是3，容量是4
	fmt.Println("newSlice", newSlice) // [20, 30 ,60]
	fmt.Println("slice", slice)       // [10, 20, 30 ,60, 50]
}

func appendNewSlice() {
	// 长度4，容量4
	slice := []int{10, 20, 30, 40}
	// 此时切片容量用完了，再追加需要扩容，此处会新加数组，长度为原数组的2倍，
	// 即 newSlice 的底层数组是新数组，新切片容量为8；
	// 而 slice 的底层数组是旧数组，二者互不影响
	newSlice := append(slice, 50)
	fmt.Println("slice", slice)       // [10, 20, 30, 40]
	fmt.Println("newSlice", newSlice) // [10, 20, 30, 40, 50]
	newSlice[0] = 100
	fmt.Println("slice", slice)       // [10, 20, 30, 40]
	fmt.Println("newSlice", newSlice) // [100, 20, 30, 40, 50]
}

// 显示设置容量
func setCapSlice() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 长度为1=3-2，容量为1=3-2  source[i:j:k] 长度=j-i 容量=k-i
	slice := source[2:3:3]
	fmt.Println("source", source, "len", len(source), "cap", cap(source)) // ["Apple", "Orange", "Plum", "Banana", "Grape"]
	fmt.Println("slice", slice, "len", len(slice), "cap", cap(slice))     // ["Plum"]
	// 超出切片容量3，需要新建数组
	slice = append(slice, "Kiwi")
	fmt.Println("source", source) // ["Apple", "Orange", "Plum", "Banana", "Grape"]
	fmt.Println("slice", slice)   // ["Plum", "Kiwi"]
}

// 合并切片
func concatSlice() {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := append(s1, s2...)
	fmt.Println("s3", s3) // [1, 2, 3, 4]
}

// 迭代切片
func iteratorSlice() {
	slice := []int{10, 20, 30, 40}
	// 与数组迭代一样，可以使用 for range + 普通 for 循环
	for index, value := range slice {
		fmt.Println(index, value)
	}
}

// 函数间传递切片
func foo(slice []int) []int {
	slice[0] = 100
	return slice
}

func transfer() {
	// 1、创建一个 slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // [1, 2, 3, 4, 5]
	// 2、调用函数，传递一个切片副本，实际上内部还是传递了对数组的指针，
	// 所以 foo 内部的操作会影响 transfer 中的 slice
	slice2 := foo(slice)
	fmt.Println("slice2", slice2) // [100, 2, 3, 4, 5]
	fmt.Println("slice", slice)   // [100, 2, 3, 4, 5]
}

func main() {
	//createSlice()
	//useSlice()
	//appendSlice()
	//appendNewSlice()
	//setCapSlice()
	//concatSlice()
	//iteratorSlice()
	transfer()
}
