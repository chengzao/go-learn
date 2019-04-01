//use: go run var.go
//变量声明 var
package main

//声明未赋值
var str0 string
var str1 string = "str1"
var str2 = "str2"

//str3 := "str3"   //错误
var str3 = "tstr3-1"

func varF() {
	//var str5 int //声明str5后需要使用,否侧报错
	str4 := "str4" //声明str4赋值str4后需要使用,否侧报错
	println(str0)  //
	println(str4)  // "str4"
}

func test() {
	str3 = "str3"
	println(str1) // "str1"
	println(str2) // "str2"
	println(str3) // "str3"
}

func utfF() {
	println('a') // 97
	println("a") // "a"
}

// iota枚举
func constF() {
	// iota
	const (
		a           = iota       //0
		b           = "abc"      //"abc"
		c, d        = iota, iota //2 2
		e           = iota       //3
		f    string = "fff"      //"fff"
	)

	println("a ==>", a)
	println("b ==>", b)
	println("c ==>", c)
	println("d ==>", d)
	println("e ==>", e)
	println("f ==>", f)
}

func runeByteF() {
	s := "Go编程"
	println("'s' 的len ==>", len(s))                //8
	println("'编'的len ==>", len(string(rune('编')))) //3
	println("rune的len ==>", len([]rune(s)))        //4
}

// iota
func init() {
	const (
		a = iota
		b
		c
	)
	println("init a b c:", a, b, c)
}

func main() {
	println("main")
	varF()
	test()
	utfF()
	constF()
	runeByteF()
}
