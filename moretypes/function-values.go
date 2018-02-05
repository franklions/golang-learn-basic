package main

import (
	"fmt"
	"math"
)

//参数是个函数
func compute(fn func(float64, float64) float64) float64{
	return fn(3,4)
}

//返回值是函数 函数的闭包
func adder() func(int) int{
	sum := 0
	return func(x int) int{
		sum +=x
		return sum
	}
}

func main() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	//
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()

	for i := 0;i<10;i++{
		fmt.Println(pos(1),neg(-2*i))
	}

}
