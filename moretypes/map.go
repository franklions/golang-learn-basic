package main

import (
	"fmt"
)

type Vertex struct{
	Lat, Long float64
}

var m map[string]Vertex

var mm = map[string]Vertex{
	"Bell Labs":Vertex{
		40.68232,-74.23023,
	},
	"Google":Vertex{
		37.232323,-122.2332,
	},
}

func mutating_maps(){
	m :=make(map[string]int)

	m["Answer"] =21

	m["Answer"] =48

	delete(m, "Answer")

	elem, ok := m["Answer"]

	fmt.Println(elem,ok)
}

func main() {

	if m ==nil {
		fmt.Println(m)
	}

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.681233, -74.232343,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(mm)
	fmt.Println(mm["Google"])

	elem, ok := mm["Google"]

	fmt.Println(elem,ok)
	mutating_maps()
}
