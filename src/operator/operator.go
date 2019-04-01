// 运算
package main

import "fmt"

func main() {
	// simpaleF()
	// equalF()
	// logicF()
	// bitF()
	// assF()
	othF()
	// testF()
}

func simpaleF() {
	var a int = 21
	var b int = 10
	var c int
	c = a + b
	fmt.Printf("a + b : %d\n", c)
	c = a - b
	fmt.Printf("a - b : %d\n", c)
	c = a * b
	fmt.Printf("a * b : %d\n", c)
	c = a / b
	fmt.Printf("a / b : %d\n", c)
	c = a % b
	fmt.Printf("c求余: %d\n", c)
	a++
	fmt.Printf("a++ : %d\n", a)
	a--
	fmt.Printf("a-- : %d\n", a)
}

func equalF() {
	var a int = 21
	var b int = 10
	fmt.Printf("a == b : %v \n", a == b)
	fmt.Printf("a > b : %v \n", a > b)
	fmt.Printf("a < b : %v \n", a < b)
	fmt.Printf("a <= b : %v \n", a <= b)
	fmt.Printf("a >= b : %v \n", a >= b)
	fmt.Printf("a != b : %v \n", a != b)
}

func logicF() {
	var a bool = true
	var b bool = false
	fmt.Printf("%v \n", a && b)
	fmt.Printf("%v \n", a || b)

	/* 修改 a 和 b 的值 */
	a = false
	b = true
	fmt.Printf("%v \n", a && b)
	fmt.Printf("%v \n", !(a && b))
}

func bitF() {
	var a uint = 60                    /* 60 = 0011 1100 */
	var b uint = 13                    /* 13 = 0000 1101 */
	fmt.Printf("a & b : %d \n", a&b)   // 12  0000 1100
	fmt.Printf("a | b : %d \n", a|b)   // 61  0011 1101
	fmt.Printf("a ^ b : %d \n", a^b)   // 49 0011 0001
	fmt.Printf("a << 2 : %d \n", a<<2) //220  1111 0000
	fmt.Printf("a >> 2 : %d \n", a>>2) //15   0000 1111
}

func assF() {
	var a int = 21
	var c int
	c = a
	fmt.Printf("c : %v \n", c)
	c += a
	fmt.Printf("c += a : %v \n", c)
	c -= a
	fmt.Printf("c -= a : %v \n", c)
	c /= a
	fmt.Printf("c /= a : %v \n", c)
	c *= a
	fmt.Printf("c *= a : %v \n", c)

	c <<= 2
	fmt.Printf("c <<= a : %v \n", c)
	c >>= 2
	fmt.Printf("c >>= a : %v \n", c)

	c &= 2
	fmt.Printf("c &= a : %v \n", c)
	c ^= 2
	fmt.Printf("c ^= a : %v \n", c)
	c |= 2
	fmt.Printf("c |= a : %v \n", c)
}

func othF() {
	var a int = 4
	var b int32
	var c float32
	var d = 321
	var ptr *int

	fmt.Printf("a类型: %T \n", a)
	fmt.Printf("b类型: %T \n", b)
	fmt.Printf("c类型: %T \n", c)
	fmt.Printf("d类型: %T \n", d)

	ptr = &a
	fmt.Printf("a值 : %d \n", a)
	fmt.Printf("*ptr : %d \n", *ptr)
}

// func testF(){
// 	var a int8 = 1
// 	var b int32 = 2
// 	var c = a + b
// 	fmt.Printf(c)  //invalid operation: a + b (mismatched types int8 and int32)
// }
