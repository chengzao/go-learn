//函数作为 值、类型
package main

import "fmt"

//type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
type testInt func(int) bool // 声明了一个函数类型

// isOdd变量类型和返回类型 与 声明的函数testInt类型一致
func isOdd(a int) bool {
	if a%2 == 0 {
		return false
	}
	return true
}

// isEven变量类型和返回类型 与 声明的函数testInt类型一致
func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}

// 声明了一个函数类型的参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		// fmt.Println(f(value))
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven) // 函数当做值来传递
	fmt.Println("Even elements of slice are: ", even)
}
