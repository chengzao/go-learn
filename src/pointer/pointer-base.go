package pointer

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	baseF()
	testF()
	pointerF()
}

// 类型*T 是 指向类型T的值 的指针.其零值是 nil
func baseF() {
	var p *int             // p 是 int类型的指针
	fmt.Printf("%v \n", p) // nil

	var i int = 6

	// & 符号会生成一个指向其作用对象的指针
	p = &i

	// * 符号表示指针指向的底层的值
	val := *p

	fmt.Printf("&i ==> %v \n", p)   // 0xxx  通过指针 p 获取 i 的地址
	fmt.Printf("*p ==> %v \n", val) // 6 通过指针 p 获得 i 的值
}

func testF() {
	var p *int
	var i int = 5
	p = &i //← 通过指针 p 获取 i 的地址
	*p = 8 //← 通过指针 p 修改 i 的值
	s := *p + 1
	fmt.Printf("%v\n", p)  // 0xc0xxx
	fmt.Printf("%v\n", *p) //← 8
	fmt.Printf("%v\n", s)  //← 9
	fmt.Printf("%v\n", i)  //← 8
}

func pointerF() {
	v := Vertex{1, 2}
	p := &v
	p.X = 123
	fmt.Printf("v ==> %v \n", v)
	fmt.Printf("p ==> %v \n", p)
	fmt.Printf("*p ==> %v \n", *p)
}
