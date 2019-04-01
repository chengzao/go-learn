package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return v.X + v.Y
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) //1.4142135623730951
	v := Vertex{3, 4}

	a = f                        //  MyFloat
	fmt.Printf("%v \n", a.Abs()) //1.4142135623730951
	a = &v                       //  *Vertex

	// 下面一行，v 是一个 Vertex(而不是 *Vertex)
	// 所以没有实现 Abser
	//a = v
	fmt.Printf("%T %v \n", a.Abs(), a.Abs()) //float64 7
}
