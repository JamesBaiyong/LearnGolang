package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // 返回一个指向结构体的指针
)

func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2} //结构体字段通过点号访问
	fmt.Println(v.X) 

	p := &Vertex{1,2} //通过指针去访问结构体
	p.X = 3
	fmt.Println(p.X)

	fmt.Println(v1,v2,v3,p)
}
