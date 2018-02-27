package main

/**
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示失败。
 */

import (
	"fmt"
	"time"
	"errors"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s",e.When,e.What)
}

func run() error{
	return &MyError{ time.Now(),"it didn't work."}
}

type ErrNegativeSqrt float64

func(fe ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(fe))
}

func Sqrt(x ErrNegativeSqrt) (ErrNegativeSqrt, error) {
	if x < 0 {
		return x,errors.New(x.Error())
	}
	return x,nil
}
func main() {
	if err :=run(); err != nil {
		fmt.Println(err)
	}

	f,err := Sqrt(-10)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(f)
}
