/**
 返回值
 */
package main

import "fmt"

/**
多值返回
 */
func swap(x, y string) (string, string){
	return y, x
}

/**
重命名返回值
 */
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	a, b :=swap("hello","world")
	fmt.Println(a,b)
	fmt.Println(split(56))
}
