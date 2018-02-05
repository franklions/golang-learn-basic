package main

import "fmt"

//定义结构体
type Vertex struct {
	X int
	Y int
}

//结构体字段使用点号来访问。
func struct_fields(){
	v := Vertex{1,2}
	v.X =4
	fmt.Println(v.X)
	fmt.Println(v)
}

func struct_pointers(){
	 v := Vertex{1, 2}
	 p := &v
	 p.X = 1e9
	 fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  	// 类型为Vertex
	v2 = Vertex{X:1}				//Y:0 被省略
	v3 = Vertex{}					//X:0  和 Y:0
	p = &Vertex{4,5}		//类型为 *Vertex
)

func main() {
	struct_fields()
	fmt.Println(Vertex{1,2})
	struct_pointers()

	fmt.Println(v1,v2,v3,p,*p)
}
