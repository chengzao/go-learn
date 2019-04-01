package main

import (
	"fmt"
	"os"
	"bufio"
)

func sum() int {
	var value int
	for i := 0; i <= 100; i++ {
		value += i
	}
	return value
}

// 等同于 while(true)
func deadLoop() {
	for {
		fmt.Println("this is a deadLoop")
	}
}

// Go 没有while，循环全部用 for，for的三个组件都可以省略
func printFile(filename string) {
	// 打开文件
	file, err := os.Open(filename)
	// 如果出错，结束进程
	if err != nil {
		panic(err)
	}
	// 获取读取器
	scanner := bufio.NewScanner(file)
	// 读取：It returns false when the scan stops, either by reaching the end of the input or an error
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main(){
	sum := sum()
	fmt.Printf("%v \n", sum)
	printFile("abc.txt")
}