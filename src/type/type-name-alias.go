package main

import "fmt"

//类型等价定义，相当于类型重命名
type name string

func (n name) len() int {
	return len(n)
}

func main() {
	var myname name = "taozs" //其实就是字符串类型
	l := []byte(myname)       //字符串转字节数组
	fmt.Println(len(l))       //字节长度
}
