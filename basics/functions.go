package main

import "fmt"

func add(x int,y int) int {
	return x + y
}

func main() {
	fmt.Println(add(18,12))
	fmt.Println(subtract(190,231))
	a, b :=swap("hello","world")
	fmt.Println(a,b)
}

func subtract(x ,y int) int {
	return x - y
}
