package main

/**
接口值
 */

 import (
 	"fmt"
 	"math"
 )

 type I interface {
 	M()
 }

 type T struct {
 	S string
 }

 func (t *T) M(){
	 if t == nil {
		 fmt.Println("<nil>")
		 return
	 }
 	fmt.Println(t.S)
 }

 type F float64
 func(f F) M(){
 	fmt.Println(f)
 }
func main() {
	var  i I
	i =&T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	//空接口
	var i interface{}
	describe(i)

	//空接口可保存任何类型的值。
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i I){
	fmt.Printf("(%v,%T)\n",i,i)
}