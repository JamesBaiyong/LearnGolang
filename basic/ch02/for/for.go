package main

import "fmt"

func go_for() {
	// go的基本for循环写法
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func go_while() {
	// 类似于while循环,修改１为０，或则去掉循环条件，程序就是死循环，类似于while True
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func go_jiujiu(){
	// 九九乘法表
	for i := 1; i<10; i++ {
		for j := 1; j<=i; j++{
			fmt.Printf("%v * %v = %v 　",i,j,i*j)
		}
		fmt.Println()
	}
}

func main(){
	go_for()
	go_while()
	go_jiujiu()
}
