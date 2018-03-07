package main

import "fmt"

func main() {
	var a [2] string
	fmt.Println(a)
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	parms := [8] int {1,2,3,4,5,6}
	fmt.Println(parms)
}
