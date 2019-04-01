package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close() // 将 "file.Close()" 压入 defer 栈中

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 将 "writer.Flush()" 压入 defer 栈中

	word := "当函数执行结束时，从 defer 栈中执行语句 - 后进先出，先 'writer.Flush()'，再 'file.Close()'"
	fmt.Fprintln(writer, word)
	// 当函数执行结束时，从 defer 栈中执行语句 - 后进先出，先 "writer.Flush()"，再 "file.Close()"
}

func deferF() {
	defer fmt.Println("this is three")
	defer fmt.Println("this is two")
	fmt.Println("this is one")

	// → one two three
}

func callMultiDefer() {
	//很多调用defer，那么defer是采用后进先出模式
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) //43210
	}
}

func main() {
	writeFile("defer.txt")
	deferF()
}
