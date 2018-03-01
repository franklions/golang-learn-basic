package main

import "fmt"

/**
通常情况下无需关闭它们。只有在必须告诉接收者不再有值需要发送的时候才有必要关闭
 */

import "fmt"

func fibonacci(n int, c chan int){
	x,y :=0,1
	for i :=0;i<n;i++{
		c <- x
		x,y = y, x+y
	}
	fmt.Println("send c:")
	close(c)
}

func main() {
	c :=make(chan int ,10)
	go fibonacci(cap(c),c)
	fmt.Println("begin receive c:")
	for i:=range c{			//这里开始阻塞
		fmt.Println(i)
	}
}
