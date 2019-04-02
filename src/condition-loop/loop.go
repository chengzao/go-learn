package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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
	// 没有条件的死循环
	for {
		fmt.Println("this is a deadLoop")
	}
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// run:  go run loop.go
func main() {
	sum := sum()
	fmt.Printf("%v \n", sum)
	printFile("abc.txt")
	fmt.Println("convertToBin results:")
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(0),
	)
}
