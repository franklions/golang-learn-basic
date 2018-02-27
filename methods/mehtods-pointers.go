package main
/**
定义指针接收者
选择值或指针作为接收者
使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
 */
 import (
 	"fmt"
 	"math"
 )

 type Vertex struct {
 	X, Y float64
 }

 func (v Vertex) Abs() float64 {
 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
 }

 //指针方法返回的是自身的引用
 func (v *Vertex) Scale(f float64){
 	v.X = v.X * f
 	v.Y = v.Y *f
 }

 /**
 把方法写成函数
  */
  func Abs(v Vertex) float64 {
  	return math.Sqrt(v.X*v.X + v.Y*v.Y)
  }

  func Scale(v *Vertex, f float64){
  	v.X = v.X * f
  	v.Y = v.Y * f
  }

func main() {
	v := Vertex{3, 4}
	var ver Vertex
	ver.Scale(20)
	fmt.Println(ver)
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())

	x := Vertex{3,4}
	Scale(&x,10)
	fmt.Println(Abs(v))
}
