package fmt

import "fmt"

func main() {
	var a int
	var b bool
	var c string
	var d float32
	var e complex64 = 5 + 2i
	fmt.Printf("a tpye is : %T \n", a)
	fmt.Printf("b tpye is : %T \n", b)
	fmt.Printf("c tpye is : %T \n", c)
	fmt.Printf("d tpye is : %T \n", d)
	fmt.Printf("e tpye is : %T \n", e)
	fmt.Printf("%+d \n", -255)
	Transtpye()
}

func Transtpye() {
	var i int16 = 2
	f := float64(i + 1)
	u := uint(f)
	fmt.Printf("i type is : %T %v \n", i, i)
	fmt.Printf("f type is : %T %v \n", f, f)
	fmt.Printf("u type is : %T %v \n", u, u)
}
