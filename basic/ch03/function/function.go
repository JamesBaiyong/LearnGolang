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

// 方法的闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 斐波纳契闭包
func fibonacci() func() int {
	value1, value2 := 0, 1

	return func() int {

		temp := value1
		
		value1,value2 = value2 , (value1+value2)
		
		return temp
	}
}


func main(){
	var (
		p Person
		t Int
	)
	p.show("Jake",23)
	p.show("James",24)
	fmt.Println(t.add(3,3))

	pos := adder()
	neg := adder()
	for i:=0;i<10;i++{
		fmt.Println(
			pos(i),
			neg(-2 * i),
		)
	}

	f := fibonacci()
	for i:=0;i<10;i++{
		fmt.Println(f())
	}
}