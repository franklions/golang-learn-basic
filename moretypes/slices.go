package main

import (
	"fmt"
	"strings"
)

func basic_slice(){
	s := []int{2,3,5,7,11,12}
	fmt.Println(s)

	for i:=0;i<len(s);i++ {
		fmt.Printf("s[%d]==%d\n",i,s[i])
	}
}

func slices_of_slice(){
	//Create a  tic-tac-toe board.
	game := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	// The player take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}
func printBoard(strs [][]string) {
	for i := 0; i < len(strs); i++ {
		fmt.Printf("%s\n", strings.Join(strs[i], " "))
	}
}

//对slice切片
func slicing_slices(){
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])
}

func making_slices(){
	a := make([]int, 5)
	printSlice("a",a)
	b := make([]int,0,5)
	printSlice("b",b)
	c := b[:2]
	printSlice("c",c)
	d := c[2:5]
	printSlice("d",d)
}

func printSlice(s string, x []int){
	fmt.Printf("%s len= %d cap=%d %v\n",s,len(x),cap(x),x)
}

func nil_slice(){
	var z []int
	fmt.Println(z,len(z),cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func append_slice(){
	var a []int
	printSlice("a",a)

	a = append(a,0)
	printSlice("a",a)

	//切片根据需要增长
	a = append(a,1)
	printSlice("a",a)    //a len= 2 cap=2 [0 1]

	a = append(a,2,3,4)
	printSlice("a",a)			//a len= 5 cap=6 [0 1 2 3 4]
}

//每次迭代 range 将返回两个值。 第一个是当前下标（序号），
// 第二个是该下标所对应元素的一个拷贝。
func range_slices(){
	pow := []int{1,2,4,8,16,32,24,128}
	for i, v := range pow{
		fmt.Printf("2**%d = %d\n",i,v)
	}

	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	//basic_slice()
	//slices_of_slice()
	//slicing_slices()
	//making_slices()
	//nil_slice()
	//append_slice()
	range_slices()
}
