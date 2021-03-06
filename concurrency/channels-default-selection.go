package main

import (
	"time"
	"fmt"
)

func main() {
	tick := time.Tick(100 * time.Microsecond)
	boom :=time.After(500 * time.Microsecond)
	for{
		select{
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("boom!")
				return
			default:
				fmt.Println("........")
				time.Sleep(50 * time.Microsecond)
			}
	}
}
