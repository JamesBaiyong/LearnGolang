package main

import (
	"fmt"
	"math"
)


type Vertex struct{
	X,Y float64
}

func (v Vertex) Abs() float64{
	// 普通的结构体类型方法
	// 这里修改的是结构体副本，本身结构体的值未改变，这是值接受者和指针接受者的差别
	v.X = v.X * 10
	v.Y = v.Y * 10
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	// 指针接收者的方法可以修改接收者指向的值
	v.X = v.X * f
	v.Y = v.Y * f
}

// 指针与函数
func Abs1(v Vertex) float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale1(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{30, 40}
	fmt.Println("比较指针接收者和非指针接收者的区别")
	fmt.Println(v)
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
	fmt.Println(v)

	fmt.Println("以下是指针与函数")

	// 指针与函数
	v1 := Vertex{3,4}
	Scale1(&v1,10)
	fmt.Println(Abs1(v1))
}
