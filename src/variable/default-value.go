package main

import "fmt"

func main() {
	normalF()
}
func normalF() {
	var (
		a1  int        //0
		a2  int8       //0
		a3  int32      //0
		a4  int64      //0
		a5  uint       //0
		a6  rune       //0 rune的实际类型是 int32
		a7  byte       //0 byte的实际类型是 uint8
		a8  float32    //0 长度为 4 byte
		a9  float64    //0 长度为 8 byte
		a10 bool       // false
		a11 string     // ""
		a12 complex64  //(0+0i)
		a13 complex128 //(0+0i)
	)
	fmt.Printf("a1 default is : %v \n", a1)
	fmt.Printf("a2 default is : %v \n", a2)
	fmt.Printf("a3 default is : %v \n", a3)
	fmt.Printf("a4 default is : %v \n", a4)
	fmt.Printf("a5 default is : %v \n", a5)
	fmt.Printf("a6 default is : %v \n", a6)
	fmt.Printf("a7 default is : %v \n", a7)
	fmt.Printf("a8 default is : %v \n", a8)
	fmt.Printf("a9 default is : %v \n", a9)
	fmt.Printf("a10 default is : %v \n", a10)
	fmt.Printf("a11 default is : %v \n", a11)
	fmt.Printf("a12 default is : %v \n", a12)
	fmt.Printf("a13 default is : %v \n", a13)
}

func init() {
	var i interface{}
	fmt.Println("i default is ", i)
	var v *int
	m := v
	fmt.Println("v default is ", v)
	fmt.Println("m default is ", m)

}
