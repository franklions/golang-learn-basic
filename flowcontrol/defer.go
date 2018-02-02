package main

import "fmt"

func defer_fun(){
	fmt.Println("world")
	return
}

func up_func(){
	defer defer_fun()
	fmt.Println("hello")
	return
}

//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
func defer_multi(){
	fmt.Println("counting")

	for i := 0; i< 10 ;i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	//defer up_func()
	//fmt.Println("main")

	defer_multi()
}
