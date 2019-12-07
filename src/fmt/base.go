package main

import (
	"fmt"
)

// Human desc
type Human struct {
	Name string
}

// 普通占位符
func normalF() {
	var people = Human{Name: "zhangsan"}

	// %v      相应值的默认格式
	fmt.Printf("%v \n", people) // {zhangsan}

	// %+v     打印结构体时，会添加字段名
	fmt.Printf("%+v \n", people) // {Name:zhangsan}

	// %#v     相应值的Go语法表示
	fmt.Printf("%#v \n", people) // main.Human{Name:"zhangsan"}

	// %T      相应值的类型的Go语法表示
	fmt.Printf("%T \n", people) // main.Human

	// %%      字面上的百分号，并非值的占位符
	fmt.Printf("%% \n") // %
}

// 布尔占位符
func booleanF() {
	// %t          true 或 false。
	fmt.Printf("%t \n", true) // true
}

// 整数占位符
func intF() {
	// %b      二进制表示
	fmt.Printf("%b \n", 5) // 101
	// %c      相应Unicode码点所表示的字符
	fmt.Printf("%c \n", 0x4E2D) // 中
	// %d      十进制表示
	fmt.Printf("%d \n", 0x12) // 18
	// %o      八进制表示
	fmt.Printf("%d \n", 10) // 12
	// %q      单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Printf("%q \n", 0x4E2D) // '中'
	// %x      十六进制表示，字母形式为小写 a-f
	fmt.Printf("%x \n", 13) // d
	// %X      十六进制表示，字母形式为大写 A-F
	fmt.Printf("%x \n", 13) // D
	// %U      Unicode格式：U+1234，等同于 "U+%04X"
	fmt.Printf("%U \n", 0x4E2D) // U+4E2D
}

// 浮点数和复数的组成部分（实部和虚部）
func floatF() {
	// %b      无小数部分的，指数为二的幂的科学计数法
	fmt.Printf("%b \n", 10.2) // 5742089524897382p-49
	// %e      科学计数法
	fmt.Printf("%e \n", 10.2) // 1.020000e+01
	// %E      科学计数法
	fmt.Printf("%e \n", 10.2) // 1.020000E+01
	// %f      有小数点而无指数
	fmt.Printf("%f \n", 10.2) // 10.200000
	// %g      根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%g \n", 10.20) // 10.2
	// %G      根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
	fmt.Printf("%G \n", 10.20+2i) // (10.2+2i)
}

// 字符串与字节切片
func stringAndSliceF() {
	// %s      输出字符串表示（string类型或[]byte)
	fmt.Printf("%s \n", []byte("Go语言")) // Go语言
	// %q      双引号围绕的字符串，由Go语法安全地转义
	fmt.Printf("%q \n", "Go语言") // "Go语言"
	// %x      十六进制，小写字母，每字节两个字符
	fmt.Printf("%x \n", "golang") // 676f6c616e67
	// %X      十六进制，大写字母，每字节两个字符
	fmt.Printf("%X \n", "golang") // 676F6C616E67
}

// 指针
func pointerF() {
	// %p      十六进制表示，前缀 0x
	people := Human{Name: "zhangsan"}
	fmt.Printf("%p \n", &people) // 0x4f57xxxxx
}

// 其它标记
func otherF() {
	// +      总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。
	fmt.Printf("%+q \n", "中文") // "\u4e2d\u6587"
	// -      在右侧而非左侧填充空格（左对齐该区域）
	// #      备用格式：为八进制添加前导 0（%#o）
	fmt.Printf("%#U \n", '中') // U+4E2D
	//  为十六进制添加前导 0x（%#x）或 0X（%#X），为 %p（%#p）去掉前导 0x；
	//  如果可能的话，%q（%#q）会打印原始 （即反引号围绕的）字符串；
	//  如果是可打印字符，%U（%#U）会写出该字符的
	//  Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
	// ' '    (空格)为数值中省略的正负号留出空白（% d）；
	//  以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
	// 0      填充前导的0而非空格；对于数字，这会将填充移到正负号之后
}

func main() {
	fmt.Println("fmt mod")
	fmt.Println("----普通占位符----")
	normalF()
	fmt.Println("----布尔占位符----")
	booleanF()
	fmt.Println("----整数占位符----")
	intF()
	fmt.Println("----浮点数和复数的组成部分（实部和虚部）----")
	floatF()
	fmt.Println("----字符串与字节切片----")
	stringAndSliceF()
	fmt.Println("----指针----")
	pointerF()
	fmt.Println("----其它标记----")
	otherF()
}
