package main

import (
	"fmt"
)

func main() {
	// sliceF2()
	// sliceAppendF()
	// sliceCopyF()
	sliceTestF()
	// sliceVarF()
	// sliceVarF2()
}

func sliceVarF() {
	var array = [...]byte{'a', 'b', 'c', 'd'}
	var slice1 []byte
	slice1 = array[:3]
	fmt.Printf("array[:3] ==> %v \n", slice1)
	fmt.Printf("array[:] ==> %v \n", array[:])
	fmt.Printf("array[0:] ==> %v \n", array[0:])
}

func sliceVarF2() {
	//声明
	var d = [10]int{2: 4, 5: 7}
	for i, v := range d {
		println(i, v)
	}
}

func sliceAppendF() {
	//使用make 创建 切片
	var slice1 = make([]int, 3, 6)
	// 使用 append 添加元素，并且未超出 cap
	slice2 := append(slice1, 1, 2, 3)
	// 使用 append 添加元素，并且超出 cap. 这个时候底层数组会变化，新增加的元素只会添加到新的底层数组，不会覆盖旧的底层数组。
	slice3 := append(slice1, 4, 5, 6, 7)
	slice1[0] = 10
	fmt.Printf("len = %d cap = %d %v\n", len(slice1), cap(slice1), slice1)
	fmt.Printf("len = %d cap = %d %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len = %d cap = %d %v\n", len(slice3), cap(slice3), slice3)
}

func sliceF2() {
	var slice1 = []int{1, 2, 3, 4, 5}
	//使用下标访问 slice
	for i := 0; i <= 4; i++ {
		fmt.Println(i, slice1[i])
	}

	//使用range 进行遍历
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
}

func sliceCopyF() {
	var slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice2 = make([]int, 3, 5)
	var n int
	n = copy(slice2, slice1) // just copy three elements
	fmt.Println(n, slice2, len(slice2), cap(slice2))

	slice3 := slice1[3:6]          //二者引用同一个底层数组
	fmt.Println(slice3)            // 456
	fmt.Println(slice1[1:5])       // 2345
	n = copy(slice3, slice1[1:5])  //所以，copy的时候发生元素重叠
	fmt.Println(n, slice1, slice3) //3 [1 2 3 2 3 4 7 8 9 10] [2 3 4]
}

func sliceTestF() {
	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[2:4]
	s2 := a[1:5]
	s3 := a[:]
	s4 := a[:4]
	s5 := s2[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println("cap", cap(a[1:3:5]))
	fmt.Println("len", len(a[1:3:5]))
}
