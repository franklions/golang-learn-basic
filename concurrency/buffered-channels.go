package main

import "fmt"

func main() {
	ch := make(chan int,2)
	ch <- 1
	ch <- 2
	//ch <- 3				//  会产生错误fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
