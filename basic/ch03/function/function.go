package main

import "fmt"


type Person struct{
	Name string
	Age int
}

func (student *Person) show(name string,age int) {
	//针对Person的方法
	student.Name = name
	student.Age = age
	fmt.Printf("Name is %s,Age is %d",student.Name,student.Age)
	fmt.Println()
}

type Int int

func (t *Int) add(i,j int) int {
	//加法
	return i + j
}

func main(){
	var p Person
	var t Int
	p.show("Jake",23)
	p.show("James",24)
	fmt.Println(t.add(3,3))
}