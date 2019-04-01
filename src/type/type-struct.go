package main

import "fmt"

//声明新的类型,作为其它类型的属性或字段的容器
type person struct {
	name string
	age  int
}

func main() {
	// P现在就是person类型的变量了
	var P person
	// 赋值"Astaxie"给P的name属性.
	P.name = "Astaxie"
	// 赋值"25"给变量P的age属性
	P.age = 25
	// 访问P的name属性.
	fmt.Printf("The person's name is %s ,age is %d \n", P.name, P.age)

	// 按照顺序提供初始化值
	P1 := person{"Tom", 25}
	fmt.Printf("%v \n", P1)

	// 通过field:value的方式初始化，这样可以任意顺序
	P2 := person{age: 24, name: "Tom"}
	fmt.Printf("%v \n", P2)

	// 通过new函数分配一个指针，此处P的类型为*person
	P3 := new(*person)
	fmt.Println(P3)
}
