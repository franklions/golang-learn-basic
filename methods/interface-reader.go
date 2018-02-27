package main

import (
	"strings"
	"fmt"
	"io"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte,8)

	for{
		n,err :=r.Read(b)
		fmt.Printf("n=%v err=%v\n",n,err,b)
		fmt.Print("b[:n]=%q\n",b[:n])

		if err == io.EOF {
			break
		}
	}
}
