//use: go run str.go
package main

import (
	"fmt"
)

// 复数
func complexF() {
	var c complex64 = 5 + 5i
	fmt.Printf("value is:%v \n", c) // "value is:(5+5i)"
}

//字符串截取
func strSliceF() {
	// 方式1
	s := "hello"
	c := []byte(s) // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c)        // 再转换回 string 类型
	fmt.Printf("%s\n", s2) // "cello"

	// 方式2
	s = "c" + s[1:]       // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s) // "cello"
}

// 字符串连接
func strConF() {
	// 单行字符串
	str1 := "Hi golang"
	// 多行字符串
	str2 := "Starting part " +
		"Ending part"
	//多行字符串
	str3 := `
	hello
	world`
	fmt.Printf("%s\n", str1)
	fmt.Printf("%s\n", str2)
	fmt.Printf("%s\n", str3)

	fmt.Println('a')
	fmt.Println("a")
	fmt.Println(`a`)
}

//
func main() {
	//complexF()
	//strSliceF()
	strConF()
}
