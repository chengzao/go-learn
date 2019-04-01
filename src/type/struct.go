package main

import "fmt"

// 声明一个新的类型
type person struct {
	name string
	age  int
}

type Vertex struct {
	X int
	Y int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age { // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func baseF() {
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化
	bob := person{age: 25, name: "Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	tbOlder, tbDiff := Older(tom, bob)
	tpOlder, tpDiff := Older(tom, paul)
	bpOlder, bpDiff := Older(bob, paul)

	fmt.Printf(" %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tbOlder.name, tbDiff)

	fmt.Printf(" %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tpOlder.name, tpDiff)

	fmt.Printf(" %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bpOlder.name, bpDiff)
}

func pointerF() {
	v := Vertex{1, 2}
	p := &v
	p.X = 123
	fmt.Printf("v ==> %v \n", v)
	fmt.Printf("p ==> %v \n", p)
	fmt.Printf("*p ==> %v \n", *p)
}

func test1() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
}

func testF() {
	var (
		v1 = Vertex{1, 2}  // 类型为 Vertex
		v2 = Vertex{X: 1}  // Y:0 被省略  使用 Name: 语法可以仅列出部分字段
		v3 = Vertex{}      // X:0 和 Y:0
		p  = &Vertex{1, 2} // 类型为 *Vertex
	)
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("%T %T", p, v1)
}

func main() {
	// baseF()
	pointerF()
	// testF()
	test1()
}
