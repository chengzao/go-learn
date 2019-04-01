package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func readFile() {
	const filename = "defer.txt1"
	// Go 函数可以返回两个值
	// func ReadFile(filename string) ([]byte, error)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		myerr := errors.New("my custom error")
		fmt.Println(myerr)
	} else {
		// contents 是 []byte, 用%s 可以打印出来
		fmt.Printf("%s", contents)
	}
	// if 语句外部可访问
	//fmt.Printf("%s", contents)
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func errF() {
	err := errors.New("emit macho dwarf: elf header corrupted \n")
	if err != nil {
		fmt.Print(err)
	}
	// var a []byte
	// fmt.Println(a != nil)
}

func main() {
	readFile()

	if err := run(); err != nil {
		fmt.Println(err)
	}
	errF()
}
