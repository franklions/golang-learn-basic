package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

//定义类型
type MyFloat float64


//在结构体类型上定义方法
func (v *Vertex) Abs() float64{
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

//在任意类型上定义方法
func (f MyFloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//正常的函数
func Abs(v Vertex)  float64 {
	return math.Sqrt(v.X*v.X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3,4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
