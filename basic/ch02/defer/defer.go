package main

import (
	"fmt"
)

func defer_test() {
	fmt.Println("World")
}

func main() {
	defer fmt.Println("done")
	fmt.Println("Hello")
	fmt.Println("倒计时...")
	for i:=0;i<=10; i++{
		defer fmt.Println(i)
	}
	fmt.Println("go...")
	defer_test()
}
